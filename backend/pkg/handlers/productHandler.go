package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/geraldbahati/ecommerce/pkg/usecases"
	"github.com/google/uuid"
	"net/http"
)

type ProductHandler struct {
	productService *usecases.ProductService
}

func NewProductHandler(productService *usecases.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	// get page and page size
	page, pageSize, err := GetPageAndPageSize(pageStr, pageSizeStr)

	products, err := h.productService.GetProducts(r.Context(), pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error fetching products from the database.")
		return
	}

	RespondWithJSON(w, http.StatusOK, products)
}

// CreateProduct creates a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// Parameters
	var params struct {
		Name          string   `json:"name"`
		Description   string   `json:"description"`
		ImageUrl      string   `json:"image_url"`
		Price         string   `json:"price"`
		Stock         int32    `json:"stock"`
		SubCategoryID string   `json:"sub_category_id"`
		Brand         string   `json:"brand"`
		Keywords      string   `json:"keywords"`
		Colours       []string `json:"colours"`
		Materials     []string `json:"materials"`
	}

	// Decoding request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
	}

	// Add product
	product, err := h.productService.AddProduct(
		r.Context(),
		params.Name,
		params.Description,
		params.ImageUrl,
		params.Price,
		params.Stock,
		params.SubCategoryID,
		params.Brand,
		params.Keywords,
		params.Colours,
		params.Materials,
	)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to add new product: %v", err))
		return
	}

	// Respond with user
	RespondWithJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Parameters
	var params struct {
		ID            string   `json:"id"`
		Name          string   `json:"name"`
		Description   string   `json:"description"`
		ImageUrl      string   `json:"image_url"`
		Price         string   `json:"price"`
		Stock         int32    `json:"stock"`
		SubCategoryId string   `json:"category_id"`
		Brand         string   `json:"brand"`
		Rating        string   `json:"rating"`
		ReviewCount   int32    `json:"review_count"`
		DiscountRate  string   `json:"discount_rate"`
		Keywords      string   `json:"keywords"`
		IsActive      bool     `json:"is_active"`
		Colours       []string `json:"colours"`
		Materials     []string `json:"materials"`
	}

	// Decoding request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	// Update product
	product, err := h.productService.UpdateProduct(
		r.Context(),
		params.ID,
		params.Name,
		params.Description,
		params.ImageUrl,
		params.Price,
		params.Stock,
		params.SubCategoryId,
		params.Brand,
		params.Rating,
		params.ReviewCount,
		params.DiscountRate,
		params.Keywords,
		params.IsActive,
		params.Colours,
		params.Materials,
	)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to update user: %v", err))
		return
	}

	// Respond with updated product
	RespondWithJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var params struct {
		ID uuid.UUID `json:"id"`
	}

	// Decoding request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	// Deleting product
	if err := h.productService.DeleteProduct(r.Context(), params.ID); err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete product with id %v: %v", params.ID.String(), err))
		return
	}

	// Respond with success message
	RespondWithSuccess(w, http.StatusOK, "Product deleted successfully.")
}

func (h *ProductHandler) GetAvailableProducts(w http.ResponseWriter, r *http.Request) {
	// Fetching available products
	availableProducts, err := h.productService.GetAvailableProducts(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching available products from database: %v", err))
		return
	}

	// Responding with available products
	RespondWithJSON(w, http.StatusOK, availableProducts)
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	var params struct {
		ID uuid.UUID `json:"id"`
	}

	// Decoding request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body", err))
		return
	}

	// Fetching particular product
	product, err := h.productService.GetProductById(r.Context(), params.ID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching product with id %v: %v", params.ID.String(), err))
		return
	}

	// Respond with particular product
	RespondWithJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	// Parameters
	categoryIdStr := r.URL.Query().Get("category_id")

	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	// get page and page size
	page, pageSize, err := GetPageAndPageSize(pageStr, pageSizeStr)

	// Fetch products based on category
	categorizedProducts, err := h.productService.GetProductsByCategory(r.Context(), categoryIdStr, pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching products based on category id %v: %v", categoryIdStr, err))
		return
	}

	// Respond with recommended products
	RespondWithJSON(w, http.StatusOK, categorizedProducts)
}

func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	// Parameters
	var params struct {
		Query sql.NullString `json:"query"`
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	// Fetch Products based on search query
	products, err := h.productService.SearchProducts(r.Context(), params.Query)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching products based on search query %v: %v", params.Query.String, err))
		return
	}

	// Respond with search query
	RespondWithJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) GetSalesTrends(w http.ResponseWriter, r *http.Request) {
	// Fetch the month's sales trend
	currentTrend, err := h.productService.GetSalesTrends(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving the months' sales trend: %v", err))
		return
	}

	// Respond with trend
	RespondWithJSON(w, http.StatusOK, currentTrend)
}

func (h *ProductHandler) GetTrendingProducts(w http.ResponseWriter, r *http.Request) {
	// Fetch trending products
	trendingProducts, err := h.productService.GetTrendingProducts(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching trending products: %v", err))
		return
	}

	// Respond with trending products
	RespondWithJSON(w, http.StatusOK, trendingProducts)
}

// GetAllColours gets all colours
func (h *ProductHandler) GetAllColours(w http.ResponseWriter, r *http.Request) {
	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	// get page and page size
	page, pageSize, err := GetPageAndPageSize(pageStr, pageSizeStr)
	// Fetch all colours
	colours, err := h.productService.GetAllColours(r.Context(), pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching all colours: %v", err))
		return
	}

	// Respond with all colours
	RespondWithJSON(w, http.StatusOK, colours)
}

// GetAllMaterials gets all materials
func (h *ProductHandler) GetAllMaterials(w http.ResponseWriter, r *http.Request) {
	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	// get page and page size
	page, pageSize, err := GetPageAndPageSize(pageStr, pageSizeStr)

	// Fetch all materials
	materials, err := h.productService.GetAllMaterials(r.Context(), pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching all materials: %v", err))
		return
	}

	// Respond with all materials
	RespondWithJSON(w, http.StatusOK, materials)
}
