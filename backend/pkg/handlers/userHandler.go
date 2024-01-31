package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/geraldbahati/ecommerce/pkg/usecases"
	"net/http"
)

type UserHandler struct {
	userService *usecases.UserService
}

func NewUserHandler(userService *usecases.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// params
	var params struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		UserRole  string `json:"user_role"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	// create user
	user, err := h.userService.CreateUser(r.Context(), params.Email,
		params.Password, params.FirstName, params.LastName, params.UserRole)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

	// respond with user
	RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	// params
	var params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	// login user
	user, err := h.userService.LoginUser(r.Context(), params.Email, params.Password)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to login user: %v", err))
		return
	}

	// respond with user
	RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// params
	var params struct {
		RefreshToken string `json:"refresh_token"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	// refresh token
	user, err := h.userService.RefreshToken(r.Context(), params.RefreshToken)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to refresh token: %v", err))
		return
	}

	// respond with user
	RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// params
	var params struct {
		Email       string `json:"email"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
		DateOfBirth string `json:"date_of_birth"`
		Gender      string `json:"gender"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	// update user
	user, err := h.userService.UpdateUser(r.Context(), params.Email, params.FirstName, params.LastName,
		params.PhoneNumber, params.Gender, params.DateOfBirth)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to update user: %v", err))
		return
	}

	// respond with user
	RespondWithJSON(w, http.StatusOK, user)
}
