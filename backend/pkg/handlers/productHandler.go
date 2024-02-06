package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/geraldbahati/ecommerce/pkg/usecases"
	"github.com/google/uuid"
)

type ProductHandler struct{
	productService *usecases.ProductService
}

func NewProductHandler(productService *usecases.ProductService) *ProductHandler{
	return &ProductHandler{
		productService: productService,
	}
}


func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request){
	products, err := h.productService.GetProducts(r.Context())
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, "Error fetching products from the database.")
		return
	}

	RespondWithJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request){
	// Parameters
	var params struct{
		ID           uuid.UUID			`json:"id"`
		Name         string				`json:"name"`
		Description  sql.NullString		`json:"description"`
		ImageUrl     sql.NullString		`json:"image_url"`
		Price        string				`json:"price"`
		Stock        int32				`json:"stock"`
		CategoryID   uuid.UUID			`json:"category_id"`
		Brand        sql.NullString		`json:"brand"`
		Rating       string				`json:"rating"`
		ReviewCount  int32				`json:"review_count"`
		DiscountRate string				`json:"discount_rate"`
		Keywords     sql.NullString		`json:"keywords"`
	}

	// Decoding request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil{
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
	}
	
	// Add product
	product, err := h.productService.AddProduct(r.Context(), params.ID, params.Name, params.Description, params.ImageUrl, params.Price, params.Stock, params.CategoryID, params.Brand, params.Rating, params.ReviewCount, params.DiscountRate, params.Keywords, true, time.Now(), sql.NullTime{})
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to add new product: %v", err))
		return
	}
	
	// Respond with user
	RespondWithJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request){
	// Parameters
	var params struct {
		ID           uuid.UUID			`json:"id"`
		Name         string				`json:"name"`
		Description  sql.NullString		`json:"description"`
		ImageUrl     sql.NullString		`json:"image_url"`
		Price        string				`json:"price"`
		Stock        int32				`json:"stock"`
		CategoryID   uuid.UUID			`json:"category_id"`
		Brand        sql.NullString		`json:"brand"`
		Rating       string				`json:"rating"`
		ReviewCount  int32				`json:"review_count"`
		DiscountRate string				`json:"discount_rate"`
		Keywords     sql.NullString		`json:"keywords"`
		IsActive	 bool				`json:"is_active"`
	}

	// Decoding request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil{
		RespondWithError(w, http.StatusBadRequest, fmt.Sprint("Failed to decode request body: %v", err))
		return
	}

	// Update product
	product, err := h.productService.UpdateProduct(r.Context(), params.ID, params.Name, params.Description, params.ImageUrl, params.Price, params.Stock, params.CategoryID, params.Brand, params.Rating, params.ReviewCount, params.DiscountRate, params.Keywords, params.IsActive)
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to update user: %v", err))
		return
	}

	// Respond with updated product
	RespondWithJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request){
	var params struct{
		ID		uuid.UUID		`json:"id"`
	}

	// Decoding request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil{
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	// Deleting product
	if err := h.productService.DeleteProduct(r.Context(), params.ID); err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete product with id %v: %v", params.ID.String(), err))
		return
	}

	// Respond with success message
	RespondWithSuccess(w, http.StatusOK, "Product deleted successfully.")
}

func (h *ProductHandler) GetAvailableProducts(w http.ResponseWriter, r *http.Request){
	// Fetching available products
	availableProducts, err := h.productService.GetAvailableProducts(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError,fmt.Sprintf("Error fetching available products from database: %v", err))
		return
	}
	
	// Responding with available products
	RespondWithJSON(w, http.StatusOK, availableProducts)
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request){
	var params struct{
		ID		uuid.UUID		`json:"id"`
	}

	// Decoding request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil{
		RespondWithError(w, http.StatusBadRequest,fmt.Sprintf("Failed to decode request body", err))
		return
	}

	// Fetching particular product
	product, err := h.productService.GetProductById(r.Context(), params.ID)
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching product with id %v: %v", params.ID.String() ,err))
		return
	}

	// Respond with particular product
	RespondWithJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) GetFilteredProducts(w http.ResponseWriter, r *http.Request){
	// Parameters
	var params struct{
		CategoryID		uuid.UUID		`json:"category_id"`
		Price			string			`json:"price"`
	}

	// Decoding the request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil{
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body:%v", err))
		return
	}

	// Fetching filtered products
	filteredProducts, err := h.productService.GetFilteredProducts(r.Context(), params.CategoryID, params.Price)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching products with filter queries; %v, %v: %v", params.CategoryID.String(), params.Price, err))
		return
	}

	// Respond with filtered products
	RespondWithJSON(w, http.StatusOK, filteredProducts)
}

func (h *ProductHandler) GetPaginatedProducts(w http.ResponseWriter, r *http.Request){
	// Parameters 
	var params struct{
		Offset		int32		`json:"offset"`
		LIMIT		int32		`json:"limit"`
	}

	// Decoding request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil{
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body:%v", err))
		return
	}

	// Fetching paginated products
	paginatedProducts, err := h.productService.GetPaginatedProducts(r.Context(), params.Offset, params.LIMIT)
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching products with applied paginations; %v, %v: %v",params.Offset, params.LIMIT, err))
		return
	}

	// Respond with paginated results
	RespondWithJSON(w, http.StatusOK, paginatedProducts)
}

func (h *ProductHandler) GetProductWithRecommendations(w http.ResponseWriter, r *http.Request){
	// Parameters
	var params struct{
		ID		uuid.UUID		`json:"id"`
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil{
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	// Fetch products based on recommendations
	recommendedProducts, err := h.productService.GetProductWithRecommendations(r.Context(), params.ID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to fetch recommended products based on product with id %v: %v", params.ID.String(), err))
		return
	}

	// Respond with recommended products
	RespondWithJSON(w, http.StatusOK, recommendedProducts)
}

func (h *ProductHandler) GetProductsByCategory(w http.ResponseWriter, r *http.Request){
	// Parameters
	var params struct{
		CategoryID		uuid.UUID		`json:"category_id"`
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil{
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	// Fetch products based on category
	categorizedProducts, err := h.productService.GetProductsByCategory(r.Context(), params.CategoryID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to fetch products based on category with id %v: %v", params.CategoryID.String(), err))
		return
	}

	// Respond with recommended products
	RespondWithJSON(w, http.StatusOK, categorizedProducts)
}

func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request){
	// Parameters
	var params struct{
		Query		sql.NullString		`json:"query"`
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil{
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

func (h *ProductHandler) GetSalesTrends(w http.ResponseWriter, r *http.Request){
	// Fetch the month's sales trend
	currentTrend, err := h.productService.GetSalesTrends(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving the months' sales trend: %v", err))
		return
	}

	// Respond with trend
	RespondWithJSON(w, http.StatusOK, currentTrend)
}