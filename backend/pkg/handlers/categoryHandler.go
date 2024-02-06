package handlers

import (
	"encoding/json"
	"github.com/geraldbahati/ecommerce/pkg/usecases"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	categoryService *usecases.CategoryService
}

func NewCategoryHandler(categoryService *usecases.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	// params
	var params struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		ImageUrl    string `json:"image_url"`
		SeoKeywords string `json:"seo_keywords"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	// create category
	category, err := h.categoryService.CreateCategory(r.Context(), params.Name, params.Description, params.ImageUrl, params.SeoKeywords)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to create category")
		return
	}

	// respond with category
	RespondWithJSON(w, http.StatusOK, category)
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	var page int32 = 0
	var pageSize int32 = 0

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = int32(p)
	}

	if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
		pageSize = int32(ps)
	}
	// get categories
	categories, err := h.categoryService.GetAllCategories(r.Context(), pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to get categories")
		return
	}

	// respond with categories
	RespondWithJSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	// get category id
	vars := mux.Vars(r)
	categoryId, err := uuid.Parse(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid category id")
		return
	}

	// get category
	category, err := h.categoryService.GetCategoryById(r.Context(), categoryId)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to get category")
		return
	}

	// respond with category
	RespondWithJSON(w, http.StatusOK, category)
}

func (h *CategoryHandler) SearchCategoriesByName(w http.ResponseWriter, r *http.Request) {
	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")
	name := r.URL.Query().Get("name")

	var page int32 = 0
	var pageSize int32 = 0

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = int32(p)
	}

	if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
		pageSize = int32(ps)
	}

	// search categories
	categories, err := h.categoryService.SearchCategoriesByName(r.Context(), name, pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to search categories")
		return
	}

	// respond with categories
	RespondWithJSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// get category id
	vars := mux.Vars(r)
	categoryId, err := uuid.Parse(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid category id")
		return
	}

	// params
	var params struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		ImageUrl    string `json:"image_url"`
		SeoKeywords string `json:"seo_keywords"`
		IsActive    bool   `json:"is_active"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to decode request body")
		return
	}

	// update category
	category, err := h.categoryService.UpdateCategory(r.Context(), categoryId, params.Name, params.Description, params.ImageUrl, params.SeoKeywords, params.IsActive)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to update category")
		return
	}

	// respond with category
	RespondWithJSON(w, http.StatusOK, category)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// get category id
	vars := mux.Vars(r)
	categoryId, err := uuid.Parse(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid category id")
		return
	}

	// delete category
	err = h.categoryService.DeleteCategory(r.Context(), categoryId)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to delete category")
		return
	}

	// respond with success
	RespondWithJSON(w, http.StatusOK, "Category deleted successfully")
}

func (h *CategoryHandler) GetActiveCategories(w http.ResponseWriter, r *http.Request) {
	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	var page int32 = 0
	var pageSize int32 = 0

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = int32(p)
	}

	if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
		pageSize = int32(ps)
	}

	// get active categories
	categories, err := h.categoryService.GetActiveCategories(r.Context(), pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to get active categories")
		return
	}

	// respond with categories
	RespondWithJSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) GetInactiveCategories(w http.ResponseWriter, r *http.Request) {
	// get page and page size
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	var page int32 = 0
	var pageSize int32 = 0

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = int32(p)
	}

	if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
		pageSize = int32(ps)
	}

	// get inactive categories
	categories, err := h.categoryService.GetInactiveCategories(r.Context(), pageSize, page)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to get inactive categories")
		return
	}

	// respond with categories
	RespondWithJSON(w, http.StatusOK, categories)
}
