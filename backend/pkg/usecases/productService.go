package usecases

import (
	"context"
	"database/sql"
	"errors"
	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/geraldbahati/ecommerce/pkg/repository"
	"github.com/geraldbahati/ecommerce/pkg/utils"
	"github.com/google/uuid"
	"log"
	"sync"
)

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// Get All Products
func (s *ProductService) GetProducts(ctx context.Context, pageSize int32, page int32) (model.PaginationResult, error) {
	// get product count
	productCount, err := s.productRepo.GetProductCount(ctx)
	if err != nil {
		return model.PaginationResult{}, err
	}

	paginatedProducts, err := utils.Paginate(
		ctx,
		productCount,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.productRepo.GetProducts(ctx, pageSize, page)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedProducts, nil
}

// Get a specific product details
func (s *ProductService) GetProductDetails(ctx context.Context, productID uuid.UUID) (database.Product, error) {
	return s.productRepo.GetProductById(ctx, productID)
}

type task struct {
	taskType  string
	value     string
	productID uuid.UUID
}

// worker function
func (s *ProductService) worker(ctx context.Context, tasks <-chan task, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()
	for t := range tasks {
		log.Printf("Processing %V: %v", t.taskType, t.value)
		switch t.taskType {
		case "color":
			err := s.processColor(ctx, t.value, t.productID)
			if err != nil {
				errChan <- err
				return
			}
		case "material":
			err := s.processMaterial(ctx, t.value, t.productID)
			if err != nil {
				errChan <- err
				return
			}
		}
	}
}

// AddProduct creates a new product
func (s *ProductService) AddProduct(
	ctx context.Context,
	Name string,
	Description string,
	ImageUrl string,
	Price string,
	Stock int32,
	SubCategoryID string,
	Brand string,
	Keywords string,
	Colours []string,
	Materials []string,
) (model.Product, error) {

	descriptionValue := sql.NullString{}
	if Description != "" {
		descriptionValue.String = Description
		descriptionValue.Valid = true
	}

	imageUrlValue := sql.NullString{}
	if ImageUrl != "" {
		imageUrlValue.String = ImageUrl
		imageUrlValue.Valid = true
	}

	brandValue := sql.NullString{}
	if Brand != "" {
		brandValue.String = Brand
		brandValue.Valid = true
	}

	keywordsValue := sql.NullString{}
	if Keywords != "" {
		keywordsValue.String = Keywords
		keywordsValue.Valid = true
	}

	// parse category id to uuid
	categoryUUID, err := uuid.Parse(SubCategoryID)
	categoryIDValue := uuid.NullUUID{
		UUID:  categoryUUID,
		Valid: true,
	}
	if err != nil {
		return model.Product{}, err
	}

	// create product
	createProduct := model.AddProductParams{
		Name:          Name,
		Description:   descriptionValue,
		ImageUrl:      imageUrlValue,
		Price:         Price,
		Stock:         Stock,
		SubCategoryID: categoryIDValue,
		Brand:         brandValue,
		Keywords:      keywordsValue,
	}

	// create product
	newProduct, err := s.productRepo.AddProduct(ctx, createProduct)
	if err != nil {
		return model.Product{}, err
	}

	// concurrently create product colours
	tasks := make(chan task, len(Colours)+len(Materials))
	errChan := make(chan error)
	var wg sync.WaitGroup

	// Start a predefined number of workers
	for i := 0; i < 5; i++ { // Number of workers
		wg.Add(1)
		go s.worker(ctx, tasks, &wg, errChan)
	}

	// Distribute tasks for colours and materials
	go func() {
		for _, colour := range Colours {
			tasks <- task{taskType: "color", value: colour, productID: newProduct.ID}
		}
		for _, material := range Materials {
			tasks <- task{taskType: "material", value: material, productID: newProduct.ID}
		}
		close(tasks)
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(errChan)

	// Check if there were any errors
	if err, ok := <-errChan; ok {
		log.Fatalf("Failed to create product colours or materials: %v", err)
		return model.Product{}, err
	}

	// return created product
	return newProduct, nil
}

// processColor processes product colour
func (s *ProductService) processColor(ctx context.Context, colour string, productID uuid.UUID) error {
	colourModel, err := s.productRepo.GetColourByHex(ctx, colour)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newColour, err := s.productRepo.CreateProductColour(ctx, colour)
			if err != nil {
				return err
			}
			_, err = s.productRepo.UpdateProductColour(ctx, productID, newColour.ID)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	_, err = s.productRepo.UpdateProductColour(ctx, productID, colourModel.ID)
	if err != nil {
		return err
	}
	return nil
}

// processMaterial processes product material
func (s *ProductService) processMaterial(ctx context.Context, material string, productID uuid.UUID) error {
	materialModel, err := s.productRepo.GetMaterialByName(ctx, material)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newMaterial, err := s.productRepo.CreateProductMaterial(ctx, material)
			if err != nil {
				return err
			}
			_, err = s.productRepo.UpdateProductMaterial(ctx, productID, newMaterial.ID)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	_, err = s.productRepo.UpdateProductMaterial(ctx, productID, materialModel.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(
	ctx context.Context,
	ID string,
	Name string,
	Description string,
	ImageUrl string,
	Price string,
	Stock int32,
	SubCategoryID string,
	Brand string,
	Rating string,
	ReviewCount int32,
	DiscountRate string,
	Keywords string,
	IsActive bool,
	Colours []string,
	Materials []string,
) (model.Product, error) {
	// parse product id to uuid
	productID, err := uuid.Parse(ID)
	if err != nil {
		log.Fatalf("Failed to parse product id: %v", err)
		return model.Product{}, err
	}

	// get the existing product
	existingProduct, err := s.productRepo.GetProductById(ctx, productID)
	if err != nil {
		log.Fatalf("Failed to get product: %v", err)
		return model.Product{}, err
	}

	// if the value is provided update otherwise use the existing value
	if Name == "" {
		Name = existingProduct.Name
	}

	if Price == "" {
		Price = existingProduct.Price
	}

	if Stock == 0 {
		Stock = existingProduct.Stock
	}

	if Rating == "" {
		Rating = existingProduct.Rating
	}

	if ReviewCount == 0 {
		ReviewCount = existingProduct.ReviewCount
	}

	if DiscountRate == "" {
		DiscountRate = existingProduct.DiscountRate
	}

	if Keywords == "" {
		Keywords = existingProduct.Keywords.String
	}

	descriptionValue := sql.NullString{}
	if Description != "" {
		descriptionValue.String = Description
		descriptionValue.Valid = true
	} else {
		descriptionValue.String = existingProduct.Description.String
		descriptionValue.Valid = existingProduct.Description.Valid
	}

	imageUrlValue := sql.NullString{}
	if ImageUrl != "" {
		imageUrlValue.String = ImageUrl
		imageUrlValue.Valid = true
	} else {
		imageUrlValue.String = existingProduct.ImageUrl.String
		imageUrlValue.Valid = existingProduct.ImageUrl.Valid
	}

	brandValue := sql.NullString{}
	if Brand != "" {
		brandValue.String = Brand
		brandValue.Valid = true
	}

	keywordsValue := sql.NullString{}
	if Keywords != "" {
		keywordsValue.String = Keywords
		keywordsValue.Valid = true
	}

	// parse subCategory id to uuid
	var subCategoryIDValue uuid.NullUUID
	if SubCategoryID != "" {
		subCategoryUUID, err := uuid.Parse(SubCategoryID)
		subCategoryIDValue = uuid.NullUUID{
			UUID:  subCategoryUUID,
			Valid: true,
		}
		if err != nil {
			log.Fatalf("Failed to parse sub category id: %v", err)
			return model.Product{}, err
		}
	} else {
		// get the existing sub category id
		subCategoryIDValue = existingProduct.SubCategoryID
	}

	// update product
	updateProduct := model.UpdateProductParams{
		ID:            productID,
		Name:          Name,
		Description:   descriptionValue,
		ImageUrl:      imageUrlValue,
		Price:         Price,
		Stock:         Stock,
		SubCategoryID: subCategoryIDValue,
		Brand:         brandValue,
		Rating:        Rating,
		ReviewCount:   ReviewCount,
		DiscountRate:  DiscountRate,
		Keywords:      keywordsValue,
		IsActive:      IsActive,
	}

	// update product
	var productId uuid.UUID
	updatedProduct, err := s.productRepo.UpdateProduct(ctx, updateProduct)
	if err != nil {
		log.Fatalf("Failed to update product: %v", err)
		return model.Product{}, err
	}

	// if nothing is updated in the above
	if updatedProduct.ID == (uuid.UUID{}) {
		// get the product
		product, err := s.productRepo.GetProductById(ctx, productID)
		if err != nil {
			log.Fatalf("Failed to get product: %v", err)
			return model.Product{}, err
		}
		productId = product.ID
	} else {
		productId = updatedProduct.ID
	}

	// concurrently update product colours
	tasks := make(chan task, len(Colours)+len(Materials))
	errChan := make(chan error)
	var wg sync.WaitGroup

	// Start a predefined number of workers
	for i := 0; i < 5; i++ { // Number of workers
		wg.Add(1)
		go s.worker(ctx, tasks, &wg, errChan)
	}

	// Distribute tasks for colours and materials
	go func() {
		for _, colour := range Colours {
			tasks <- task{taskType: "color", value: colour, productID: productId}
		}
		for _, material := range Materials {
			tasks <- task{taskType: "material", value: material, productID: productId}
		}
		close(tasks)
	}()

	// Wait for all workers to finish
	wg.Wait()
	close(errChan)

	// Check if there were any errors
	if err, ok := <-errChan; ok {
		log.Fatalf("Failed to create product colours or materials: %v", err)
		return model.Product{}, err
	}

	// return updated product
	return updatedProduct, nil
}

// Delete and existing product
func (s *ProductService) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	return s.productRepo.DeleteProduct(ctx, productID)
}

// Fetches a particular product
func (s *ProductService) GetProductById(ctx context.Context, id uuid.UUID) (database.Product, error) {
	return s.productRepo.GetProductById(ctx, id)
}

// Fetches available products
func (s *ProductService) GetAvailableProducts(ctx context.Context) ([]database.Product, error) {
	return s.productRepo.GetAvailableProducts(ctx)
}

// Filters products based by category
func (s *ProductService) GetProductsByCategory(ctx context.Context, categoryIdStr string, pageSize int32, page int32) (model.PaginationResult, error) {
	// parse category id to uuid
	categoryID, err := uuid.Parse(categoryIdStr)
	if err != nil {
		log.Fatalf("Failed to parse category id: %v", err)
		return model.PaginationResult{}, err
	}

	// get product count by category
	productCount, err := s.productRepo.GetProductCountByCategory(ctx, categoryID)
	if err != nil {
		return model.PaginationResult{}, err
	}

	paginatedProducts, err := utils.Paginate(
		ctx,
		productCount,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.productRepo.GetProductsByCategory(ctx, categoryID)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedProducts, nil
}

// Searches for a particular product
func (s *ProductService) SearchProducts(ctx context.Context, query sql.NullString) ([]database.Product, error) {
	return s.productRepo.SearchProducts(ctx, query)
}

// Returns a sales trend for the current month
func (s *ProductService) GetSalesTrends(ctx context.Context) ([]database.GetSalesTrendsRow, error) {
	return s.productRepo.GetSalesTrends(ctx)
}

// GetTrendingProducts Returns trending products
func (s *ProductService) GetTrendingProducts(ctx context.Context) ([]model.TrendingProduct, error) {
	trendingProducts, err := s.productRepo.GetTrendingProducts(ctx)
	if err != nil {
		return nil, err
	}
	return trendingProducts, nil
}

// GetAllColours Returns all available colours
func (s *ProductService) GetAllColours(ctx context.Context, pageSize int32, page int32) (model.PaginationResult, error) {
	// get colour count
	count, err := s.productRepo.GetColourCount(ctx)
	if err != nil {
		return model.PaginationResult{}, err
	}

	// get all colours
	paginatedColours, err := utils.Paginate(
		ctx,
		count,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.productRepo.GetAllColours(ctx, offset, limit)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedColours, nil
}

// GetAllMaterials Returns all available materials
func (s *ProductService) GetAllMaterials(ctx context.Context, pageSize int32, page int32) (model.PaginationResult, error) {
	// get material count
	count, err := s.productRepo.GetMaterialCount(ctx)
	if err != nil {
		return model.PaginationResult{}, err
	}

	// get all materials
	paginatedMaterials, err := utils.Paginate(
		ctx,
		count,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.productRepo.GetAllMaterials(ctx, offset, limit)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedMaterials, nil
}

// GetProductMaterials Returns product materials
func (s *ProductService) GetProductMaterials(ctx context.Context, productId uuid.UUID, pageSize int32, page int32) (model.PaginationResult, error) {
	// get product material count
	count, err := s.productRepo.GetMaterialCount(ctx)
	if err != nil {
		return model.PaginationResult{}, err
	}

	// get product materials
	paginatedMaterials, err := utils.Paginate(
		ctx,
		count,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.productRepo.GetProductMaterials(ctx, productId, offset, limit)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedMaterials, nil
}

// GetProductColours Returns product colours
func (s *ProductService) GetProductColours(ctx context.Context, productId uuid.UUID, pageSize int32, page int32) (model.PaginationResult, error) {
	// get product colour count
	count, err := s.productRepo.GetColourCount(ctx)
	if err != nil {
		return model.PaginationResult{}, err
	}

	// get product colours
	paginatedColours, err := utils.Paginate(
		ctx,
		count,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.productRepo.GetProductColours(ctx, productId, offset, limit)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedColours, nil
}
