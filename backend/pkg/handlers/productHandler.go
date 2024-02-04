package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/geraldbahati/ecommerce/pkg/usecases"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	// Retrieve and return a list of products
	products, err := h.productService.GetProducts(r.Context())
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve product list: %v", err))
		return
	}

	// Respond with product list
	RespondWithJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) GetProductDetails(w http.ResponseWriter, r *http.Request){
	// Extract product ID  from request parameter
	productID, err := extractProductIDFromRequest(r)

	// TODO: To be looked at on how to handle error at this point
	if err != nil {
		log.Println("Error extracting product ID from request: ", err)
		return
	}

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


// Helper functions
func extractProductIDFromRequest(r *http.Request) (uuid.UUID, error) {
    vars := mux.Vars(r)
    productID, err := uuid.Parse(vars["productID"])
    if err != nil {
        return uuid.Nil, err
    }
    return productID, nil
}

func extractUserIDFromRequest(r *http.Request) (uuid.UUID, error) {
    vars := mux.Vars(r)
    userID, err := uuid.Parse(vars["userID"])
    if err != nil {
        return uuid.Nil, err
    }
    return userID, nil
}