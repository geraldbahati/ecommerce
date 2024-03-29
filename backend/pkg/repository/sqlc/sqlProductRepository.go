package sqlc

import (
	"context"
	"database/sql"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"log"

	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/google/uuid"
)

type SQLProductRepository struct {
	DB *database.Queries
}

func NewSQLProductRepository(db *database.Queries) *SQLProductRepository {
	return &SQLProductRepository{
		DB: db,
	}
}

// AddProduct creates a new product in the database
func (r *SQLProductRepository) AddProduct(ctx context.Context, product model.AddProductParams) (model.Product, error) {
	// Add product into database
	addProduct, err := r.DB.CreateProduct(ctx, database.CreateProductParams{
		ID:            uuid.New(),
		Name:          product.Name,
		Description:   product.Description,
		ImageUrl:      product.ImageUrl,
		Price:         product.Price,
		Stock:         product.Stock,
		SubCategoryID: product.SubCategoryID,
		Brand:         product.Brand,
		Keywords:      product.Keywords,
	})
	if err != nil {
		return model.Product{}, err
	}

	// Return newly added product
	return model.Product{
		ID:            addProduct.ID,
		Name:          addProduct.Name,
		Description:   addProduct.Description,
		ImageUrl:      addProduct.ImageUrl,
		Price:         addProduct.Price,
		Stock:         addProduct.Stock,
		SubCategoryID: addProduct.SubCategoryID,
		Brand:         addProduct.Brand,
		Rating:        addProduct.Rating,
		ReviewCount:   addProduct.ReviewCount,
		DiscountRate:  addProduct.DiscountRate,
		Keywords:      addProduct.Keywords,
		IsActive:      addProduct.IsActive,
		CreatedAt:     addProduct.CreatedAt,
		LastUpdated:   addProduct.LastUpdated,
	}, err
}

// UpdateProduct updates an already existing product in the database
func (r *SQLProductRepository) UpdateProduct(ctx context.Context, product model.UpdateProductParams) (model.Product, error) {
	// Update product in the database
	log.Printf("Updating product with id %s", product.ID.String())
	updatedProduct, err := r.DB.UpdateProduct(ctx, database.UpdateProductParams{
		ID:            product.ID,
		Name:          product.Name,
		Description:   product.Description,
		ImageUrl:      product.ImageUrl,
		Price:         product.Price,
		Stock:         product.Stock,
		SubCategoryID: product.SubCategoryID,
		Brand:         product.Brand,
		Rating:        product.Rating,
		ReviewCount:   product.ReviewCount,
		DiscountRate:  product.DiscountRate,
		Keywords:      product.Keywords,
		IsActive:      product.IsActive,
	})

	if err != nil {
		log.Fatalf("Error updating product with id %s: %s", product.ID.String(), err.Error())
		return model.Product{}, err
	}

	// Return updated Product
	return model.Product{
		ID:            updatedProduct.ID,
		Name:          updatedProduct.Name,
		Description:   updatedProduct.Description,
		ImageUrl:      updatedProduct.ImageUrl,
		Price:         updatedProduct.Price,
		Stock:         updatedProduct.Stock,
		SubCategoryID: updatedProduct.SubCategoryID,
		Brand:         updatedProduct.Brand,
		Rating:        updatedProduct.Rating,
		ReviewCount:   updatedProduct.ReviewCount,
		DiscountRate:  updatedProduct.DiscountRate,
		Keywords:      updatedProduct.Keywords,
		IsActive:      updatedProduct.IsActive,
		CreatedAt:     updatedProduct.CreatedAt,
		LastUpdated:   updatedProduct.LastUpdated,
	}, err
}

// DeleteProduct implements repository.ProductRepository.
func (r *SQLProductRepository) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	deletedProduct, err := r.DB.GetProductById(ctx, productID)
	if err != nil {
		log.Printf("Error fetching product with id %s: %s", productID.String(), err.Error())
		return err
	}

	err = r.DB.DeleteProduct(ctx, deletedProduct.ID)
	if err != nil {
		log.Printf("Error deleting product with id %s: %s", productID.String(), err.Error())
		return err
	}

	return nil
}

// GetAvailableProducts implements repository.ProductRepository.
func (r *SQLProductRepository) GetAvailableProducts(ctx context.Context) ([]database.Product, error) {
	availableProducts, err := r.DB.GetAvailableProducts(ctx)
	if err != nil {
		log.Printf("Error fetching available products : %s", err.Error())
		return []database.Product{}, err
	}

	return availableProducts, nil
}

// GetProductById implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductById(ctx context.Context, id uuid.UUID) (database.Product, error) {
	product, err := r.DB.GetProductById(ctx, id)
	if err != nil {
		log.Printf("Error fetching product with id %s: %s", id.String(), err.Error())
		return database.Product{}, err
	}
	return product, nil
}

// GetProducts implements repository.ProductRepository.
func (r *SQLProductRepository) GetProducts(ctx context.Context, offset int32, limit int32) (interface{}, error) {
	products, err := r.DB.GetProducts(ctx, database.GetProductsParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		log.Printf("Error fetching all products in the database : %s", err.Error())
		return []database.Product{}, err
	}
	return products, nil
}

// GetProductsByCategory implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductsByCategory(ctx context.Context, categoryID uuid.UUID) (interface{}, error) {
	categorizedProducts, err := r.DB.GetProductsByCategory(ctx, database.GetProductsByCategoryParams{
		ID: categoryID,
	})
	if err != nil {
		log.Printf("Error fetching categorized products with category id %s: %s", categoryID.String(), err.Error())
		return []database.Product{}, err
	}
	return categorizedProducts, nil
}

// GetSalesTrends implements repository.ProductRepository.
func (r *SQLProductRepository) GetSalesTrends(ctx context.Context) ([]database.GetSalesTrendsRow, error) {
	salesTrendRow, err := r.DB.GetSalesTrends(ctx)
	if err != nil {
		log.Printf("Error fetching the row of sales trends : %s", err.Error())
		return []database.GetSalesTrendsRow{}, err
	}
	return salesTrendRow, nil
}

// SearchProducts implements repository.ProductRepository.
func (r *SQLProductRepository) SearchProducts(ctx context.Context, query sql.NullString) ([]database.Product, error) {
	queryResults, err := r.DB.SearchProducts(ctx, query)
	if err != nil {
		log.Printf("Error fetching products with query  %s: %s", query.String, err.Error())
		return []database.Product{}, err
	}
	return queryResults, nil
}

// GetTrendingProducts implements repository.ProductRepository.
func (r *SQLProductRepository) GetTrendingProducts(ctx context.Context) ([]model.TrendingProduct, error) {
	trendingProducts, err := r.DB.GetTrendingProducts(ctx)
	if err != nil {
		log.Printf("Error fetching trending products : %s", err.Error())
		return []model.TrendingProduct{}, err
	}

	// Return trending products
	var modelTrendingProducts []model.TrendingProduct
	for _, product := range trendingProducts {
		modelTrendingProducts = append(modelTrendingProducts, model.TrendingProduct{
			ProductID:    product.ProductID,
			ProductName:  product.ProductName,
			Price:        product.Price,
			CategoryID:   product.CategoryID,
			CategoryName: product.CategoryName,
			SalesVolume:  product.SalesVolume,
		})
	}

	return modelTrendingProducts, nil
}

// GetProductCountByCategory implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductCountByCategory(ctx context.Context, categoryID uuid.UUID) (int64, error) {
	productCount, err := r.DB.GetProductCountByCategory(ctx, categoryID)
	if err != nil {
		log.Printf("Error fetching product count by category id %s: %s", categoryID.String(), err.Error())
		return 0, err
	}
	return productCount, nil
}

// GetProductCount implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductCount(ctx context.Context) (int64, error) {
	productCount, err := r.DB.GetProductCount(ctx)
	if err != nil {
		log.Printf("Error fetching product count : %s", err.Error())
		return 0, err
	}
	return productCount, nil
}

// CreateProductColour creates a new product colour in the database
func (r *SQLProductRepository) CreateProductColour(ctx context.Context, colourHex string) (model.Colour, error) {
	// Add product colour into database
	log.Printf("Adding product colour with hex %s", colourHex)

	addProductColour, err := r.DB.CreateColour(ctx, database.CreateColourParams{
		ID:        uuid.New(),
		ColourHex: colourHex,
	})
	if err != nil {
		log.Fatalf("Error adding product colour with hex %s: %s\n", colourHex, err.Error())
		return model.Colour{}, err
	}

	// Return newly added product colour
	return model.Colour{
		ID:          addProductColour.ID,
		ColourHex:   addProductColour.ColourHex,
		CreatedAt:   addProductColour.CreatedAt,
		LastUpdated: addProductColour.LastUpdated,
	}, err
}

// CreateProductMaterial creates a new product material in the database
func (r *SQLProductRepository) CreateProductMaterial(ctx context.Context, materialName string) (model.Material, error) {
	// Add product material into database
	log.Printf("Adding product material with name %s", materialName)
	addProductMaterial, err := r.DB.CreateMaterial(ctx, database.CreateMaterialParams{
		ID:   uuid.New(),
		Name: materialName,
	})
	if err != nil {
		log.Fatalf("Error adding product material with name %s: %s\n", materialName, err.Error())
		return model.Material{}, err
	}

	// Return newly added product material
	return model.Material{
		ID:          addProductMaterial.ID,
		Name:        addProductMaterial.Name,
		CreatedAt:   addProductMaterial.CreatedAt,
		LastUpdated: addProductMaterial.LastUpdated,
	}, err
}

// UpdateProductColour updates an already existing product colour in the database
func (r *SQLProductRepository) UpdateProductColour(ctx context.Context, productId uuid.UUID, colourId uuid.UUID) (model.ProductColour, error) {
	// Update product colour in the database
	log.Printf("Updating product colour with id %s", colourId.String())
	updateProductColour, err := r.DB.UpdateProductColour(ctx, database.UpdateProductColourParams{
		ID:        uuid.New(),
		ProductID: productId,
		ColourID:  colourId,
	})
	if err != nil {
		log.Fatalf("Error updating product colour with id %s: %s\n", colourId.String(), err.Error())
		return model.ProductColour{}, err
	}

	// Return updated product colour
	return model.ProductColour{
		ID:          updateProductColour.ID,
		CreatedAt:   updateProductColour.CreatedAt,
		LastUpdated: updateProductColour.LastUpdated,
	}, err
}

// UpdateProductMaterial updates an already existing product material in the database
func (r *SQLProductRepository) UpdateProductMaterial(ctx context.Context, productId uuid.UUID, materialId uuid.UUID) (model.ProductMaterial, error) {
	// Update product material in the database
	log.Printf("Updating product material with id %s", materialId.String())
	updateProductMaterial, err := r.DB.UpdateProductMaterial(ctx, database.UpdateProductMaterialParams{
		ID:         uuid.New(),
		ProductID:  productId,
		MaterialID: materialId,
	})
	if err != nil {
		log.Fatalf("Error updating product material with id %s: %s\n", materialId.String(), err.Error())
		return model.ProductMaterial{}, err
	}

	// Return updated product material
	return model.ProductMaterial{
		ID:          updateProductMaterial.ID,
		CreatedAt:   updateProductMaterial.CreatedAt,
		LastUpdated: updateProductMaterial.LastUpdated,
	}, err
}

// GetAllColours gets all colours from the database
func (r *SQLProductRepository) GetAllColours(ctx context.Context, offset int32, limit int32) ([]model.Colour, error) {
	colours, err := r.DB.GetColours(ctx, database.GetColoursParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		log.Printf("Error fetching product colours : %s", err.Error())
		return []model.Colour{}, err
	}

	// Return product colours
	var modelColours []model.Colour
	for _, colour := range colours {
		modelColours = append(modelColours, model.Colour{
			ID:          colour.ID,
			ColourHex:   colour.ColourHex,
			CreatedAt:   colour.CreatedAt,
			LastUpdated: colour.LastUpdated,
		})
	}

	return modelColours, nil
}

// GetAllMaterials gets all materials from the database
func (r *SQLProductRepository) GetAllMaterials(ctx context.Context, offset int32, limit int32) ([]model.Material, error) {
	materials, err := r.DB.GetMaterials(ctx, database.GetMaterialsParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		log.Printf("Error fetching product materials : %s", err.Error())
		return []model.Material{}, err
	}

	// Return product materials
	var modelMaterials []model.Material
	for _, material := range materials {
		modelMaterials = append(modelMaterials, model.Material{
			ID:          material.ID,
			Name:        material.Name,
			CreatedAt:   material.CreatedAt,
			LastUpdated: material.LastUpdated,
		})
	}

	return modelMaterials, nil
}

// GetProductColours gets all colours of a product from the database
func (r *SQLProductRepository) GetProductColours(ctx context.Context, productId uuid.UUID, offset int32, limit int32) ([]model.Colour, error) {
	productColours, err := r.DB.GetProductColours(ctx, database.GetProductColoursParams{
		ProductID: productId,
		Offset:    offset,
		Limit:     limit,
	})
	if err != nil {
		log.Printf("Error fetching product colours with product id %s: %s", productId.String(), err.Error())
		return []model.Colour{}, err
	}

	// Return product colours
	var modelColours []model.Colour
	for _, colour := range productColours {
		modelColours = append(modelColours, model.Colour{
			ID:          colour.ID,
			ColourHex:   colour.ColourHex,
			CreatedAt:   colour.CreatedAt,
			LastUpdated: colour.LastUpdated,
		})
	}

	return modelColours, nil
}

// GetProductMaterials gets all materials of a product from the database
func (r *SQLProductRepository) GetProductMaterials(ctx context.Context, productId uuid.UUID, offset int32, limit int32) ([]model.Material, error) {
	productMaterials, err := r.DB.GetProductMaterials(ctx, database.GetProductMaterialsParams{
		ProductID: productId,
		Offset:    offset,
		Limit:     limit,
	})
	if err != nil {
		log.Printf("Error fetching product materials with product id %s: %s", productId.String(), err.Error())
		return []model.Material{}, err
	}

	// Return product materials
	var modelMaterials []model.Material
	for _, material := range productMaterials {
		modelMaterials = append(modelMaterials, model.Material{
			ID:          material.ID,
			Name:        material.Name,
			CreatedAt:   material.CreatedAt,
			LastUpdated: material.LastUpdated,
		})
	}

	return modelMaterials, nil
}

// GetColourByHex gets a colour by its hex value from the database
func (r *SQLProductRepository) GetColourByHex(ctx context.Context, hex string) (model.Colour, error) {
	colour, err := r.DB.GetColourByHex(ctx, hex)
	if err != nil {
		log.Printf("Error fetching colour with hex %s: %s", hex, err.Error())
		return model.Colour{}, err
	}

	// Return product colour
	return model.Colour{
		ID:          colour.ID,
		ColourHex:   colour.ColourHex,
		CreatedAt:   colour.CreatedAt,
		LastUpdated: colour.LastUpdated,
	}, nil
}

// GetColourCount gets the count of all colours in the database
func (r *SQLProductRepository) GetColourCount(ctx context.Context) (int64, error) {
	colourCount, err := r.DB.GetColourCount(ctx)
	if err != nil {
		log.Printf("Error fetching colour count : %s", err.Error())
		return 0, err
	}
	return colourCount, nil
}

// GetMaterialCount gets the count of all materials in the database
func (r *SQLProductRepository) GetMaterialCount(ctx context.Context) (int64, error) {
	materialCount, err := r.DB.GetMaterialCount(ctx)
	if err != nil {
		log.Printf("Error fetching material count : %s", err.Error())
		return 0, err
	}
	return materialCount, nil
}

// GetMaterialByName gets a material by its name from the database
func (r *SQLProductRepository) GetMaterialByName(ctx context.Context, materialName string) (model.Material, error) {
	material, err := r.DB.GetMaterialByName(ctx, materialName)
	if err != nil {
		log.Printf("Error fetching material with name %s: %s", materialName, err.Error())
		return model.Material{}, err
	}

	// Return product material
	return model.Material{
		ID:          material.ID,
		Name:        material.Name,
		CreatedAt:   material.CreatedAt,
		LastUpdated: material.LastUpdated,
	}, nil
}
