package handlers

import (
	"encoding/json"
	"github.com/geraldbahati/ecommerce/pkg/usecases"
	"github.com/gorilla/mux"
	"net/http"
)

type SubCategoryHandler struct {
	subCategoryService *usecases.SubCategoryService
}

func NewSubCategoryHandler(subCategoryService *usecases.SubCategoryService) *SubCategoryHandler {
	return &SubCategoryHandler{
		subCategoryService: subCategoryService,
	}
}

// CreateSubCategory creates a new sub category
func (h *SubCategoryHandler) CreateSubCategory(w http.ResponseWriter, r *http.Request) {
	// params
	var params struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		CategoryId  string `json:"category_id"`
		ImageUrl    string `json:"image_url"`
		SeoKeywords string `json:"seo_keywords"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	// create sub category
	subCategory, err := h.subCategoryService.CreateSubCategory(r.Context(), params.CategoryId, params.Name, params.Description, params.ImageUrl, params.SeoKeywords)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to create sub category")
		return
	}

	// respond with sub category
	RespondWithJSON(w, http.StatusOK, subCategory)
}

// GetProductsBySubCategory gets products by sub category
func (h *SubCategoryHandler) GetProductsBySubCategory(w http.ResponseWriter, r *http.Request) {
	// get sub category id
	subCategoryId := mux.Vars(r)["subCategoryId"]

	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	// get page and page size
	page, pageSize, err := GetPageAndPageSize(pageStr, pageSizeStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid page or page size")
		return
	}

	// get products by sub category
	products, err := h.subCategoryService.GetProductsBySubCategory(r.Context(), subCategoryId, pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to get products by sub category")
		return
	}

	// respond with products
	RespondWithJSON(w, http.StatusOK, products)
}

// ListSubCategoriesByCategory lists sub categories by category
func (h *SubCategoryHandler) ListSubCategoriesByCategory(w http.ResponseWriter, r *http.Request) {
	// get category id
	categoryId := mux.Vars(r)["categoryId"]

	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	// get page and page size
	page, pageSize, err := GetPageAndPageSize(pageStr, pageSizeStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid page or page size")
		return
	}

	// get sub categories by category
	subCategories, err := h.subCategoryService.ListSubCategoriesByCategory(r.Context(), categoryId, pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to get sub categories by category")
		return
	}

	// respond with sub categories
	RespondWithJSON(w, http.StatusOK, subCategories)
}
