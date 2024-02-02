package handlers

import (
	"fmt"
	"net/http"

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

func (h *ProductHandler) GetProductList(w http.ResponseWriter, r *http.Request){
	// Retrieve and return a list of products
	products, err := h.productService.GetProductList(r.Context())
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve product list: %v", err))
		return
	}

	// Respond with product list
	RespondWithJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) GetProductDetails(w http.ResponseWriter, r *http.Request){
	// Extract product ID  from request parameter
	productID := extractProductIDFromRequest(r)
	// TODO: To be looked at

	if productID == uuid.Nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	// Retrieve and return product details
	productDetails, err := h.productService.GetProductDetails(r.Context(), productID)
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("failed to retrieve product details: %v", err))
		return
	}

	// Respond with product details
	RespondWithJSON(w, http.StatusOK, productDetails)
}

func (h *ProductHandler) AddToWishlist(w http.ResponseWriter, r *http.Request){
	// Extract product ID from request parameter
	productID := extractProductIDFromRequest(r)
	userID := extractUserIDFromRequest(r)
	if productID == uuid.Nil{
		RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if userID == uuid.Nil{
		RespondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	// Add the product to the user's wishlist
	// TODO: Check implemetation of wishlist service (parameters)
	err := h.productService.AddToWishlist(r.Context(), userID, productID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to add product to wishlist: %v", err))
		return
	} 

	// Respond with success message
	RespondWithJSON(w, http.StatusOK, map[string]string{"message":"Product added to wishlist"})
}


// Helper functions
func extractProductIDFromRequest(r *http.Request) uuid.UUID{
	return uuid.Nil
}

func extractUserIDFromRequest(r *http.Request) uuid.UUID{
	return uuid.Nil
}