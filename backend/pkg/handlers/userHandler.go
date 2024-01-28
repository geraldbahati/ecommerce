package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/geraldbahati/ecommerce/pkg/usecases"
	"net/http"
)

type UserHandler struct {
	userService usecases.UserService
}

func NewUserHandler(userService usecases.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// params
	var params struct {
		Username  string `json:"username"`
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
	user, err := h.userService.CreateUser(r.Context(), params.Username, params.Email,
		params.Password, params.FirstName, params.LastName, params.UserRole)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

	// respond with user
	RespondWithJSON(w, http.StatusOK, user)
}
