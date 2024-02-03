package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/geraldbahati/ecommerce/pkg/usecases"
	"html/template"
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

func (h *UserHandler) UpdateProfilePicture(w http.ResponseWriter, r *http.Request) {
	// params
	var params struct {
		ProfilePicture string `json:"profile_picture"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	// update user
	user, err := h.userService.UpdateProfilePicture(r.Context(), params.ProfilePicture)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to update profile picture: %v", err))
		return
	}

	// respond with user
	RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	// params
	var params struct {
		Email string `json:"email"`
	}

	// decode request body
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to decode request body: %v", err))
		return
	}

	// send reset password email
	if err := h.userService.SendResetPasswordEmail(r.Context(), params.Email); err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to send reset password email: %v", err))
		return
	}

	// respond with success message
	RespondWithSuccess(w, http.StatusOK, "Reset password email sent successfully")
}

func (h *UserHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Serve the reset password form with the token embedded as a hidden input
		token := r.URL.Query().Get("token")
		if token == "" {
			RespondWithError(w, http.StatusBadRequest, "Token is required")
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl := template.Must(template.ParseFiles("pkg/templates/reset-password.html"))
		data := map[string]interface{}{
			"Token": token,
		}
		if err := tmpl.Execute(w, data); err != nil {
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to open reset password page: %v", err))
			return
		}
	} else if r.Method == http.MethodPost {
		// Handle form submission
		err := r.ParseForm()
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid form data")
			return
		}
		token := r.FormValue("token")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")

		// Validate the inputs
		if token == "" {
			RespondWithError(w, http.StatusBadRequest, "Token is required")
			return
		}
		if password != confirmPassword {
			RespondWithError(w, http.StatusBadRequest, "Passwords do not match")
			return
		}

		// Reset password logic...
		if err := h.userService.ResetPassword(r.Context(), token, password); err != nil {
			RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to reset password: %v", err))
			return
		}

		// Respond with success message
		RespondWithSuccess(w, http.StatusOK, "Password reset successfully")
	}
}
