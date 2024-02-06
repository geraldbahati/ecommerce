package main

import (
	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/handlers"
	"github.com/geraldbahati/ecommerce/pkg/middleware"
	"github.com/geraldbahati/ecommerce/pkg/repository/sqlc"
	"github.com/geraldbahati/ecommerce/pkg/usecases"
	"log"
	"net/http"

	"github.com/geraldbahati/ecommerce/pkg/config"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	// initialize database connection
	conn, err := config.NewDatabaseConnection(cfg.DbUrl)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
	}

	db := database.New(conn)

	// initialize repositories
	userRepo := sqlc.NewSQLUserRepository(db)
	categoryRepo := sqlc.NewSQLCategoryRepository(db)

	// initialize services
	userService := usecases.NewUserService(userRepo)
	categoryService := usecases.NewCategoryService(categoryRepo)

	// initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// setup routes
	r := mux.NewRouter()
	r.Use(middleware.CORS)

	getUserRouter(r, userHandler)
	getCategoryRouter(r, categoryHandler)

	// start server
	log.Printf("Server listening on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}

func getUserRouter(r *mux.Router, userHandler *handlers.UserHandler) {
	resetPasswordRouter := r.PathPrefix("/reset-password").Subrouter()
	resetPasswordRouter.HandleFunc("", userHandler.ResetPassword).Methods(http.MethodGet)
	resetPasswordRouter.HandleFunc("", userHandler.ResetPassword).Methods(http.MethodPost)

	userRouter := r.PathPrefix("/api/users").Subrouter()
	userRouter.HandleFunc("/register", userHandler.RegisterUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/login", userHandler.LoginUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/refresh", userHandler.RefreshToken).Methods(http.MethodPost)

	protectedUserRouter := userRouter.PathPrefix("").Subrouter()
	protectedUserRouter.Use(middleware.Auth)
	protectedUserRouter.HandleFunc("/update", userHandler.UpdateUser).Methods(http.MethodPut)
	protectedUserRouter.HandleFunc("/update-profile-picture", userHandler.UpdateProfilePicture).Methods(http.MethodPut)
	protectedUserRouter.HandleFunc("/reset-password", userHandler.RequestPasswordReset).Methods(http.MethodPut)
}

func getCategoryRouter(r *mux.Router, categoryHandler *handlers.CategoryHandler) {
	categoryRouter := r.PathPrefix("/api/categories").Subrouter()
	categoryRouter.HandleFunc("", categoryHandler.GetAllCategories).Methods(http.MethodGet)
	categoryRouter.HandleFunc("/{id}/", categoryHandler.GetCategoryById).Methods(http.MethodGet)
	categoryRouter.HandleFunc("/search", categoryHandler.SearchCategoriesByName).Methods(http.MethodGet)

	protectedCategoryRouter := categoryRouter.PathPrefix("").Subrouter()
	protectedCategoryRouter.Use(middleware.Auth)
	//protectedCategoryRouter.HandleFunc("/active", categoryHandler.GetActiveCategories).Methods(http.MethodGet)
	//protectedCategoryRouter.HandleFunc("/inactive", categoryHandler.GetInactiveCategories).Methods(http.MethodGet)
	//protectedCategoryRouter.HandleFunc("/{id}", categoryHandler.UpdateCategory).Methods(http.MethodPut)
	//protectedCategoryRouter.HandleFunc("/{id}", categoryHandler.DeleteCategory).Methods(http.MethodDelete)
	protectedCategoryRouter.HandleFunc("", categoryHandler.CreateCategory).Methods(http.MethodPost)
}
