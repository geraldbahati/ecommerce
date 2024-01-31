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

	// initialize services
	userService := usecases.NewUserService(userRepo)

	// initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	// setup routes
	r := mux.NewRouter()
	r.Use(middleware.CORS)

	getUserRouter(r, userHandler)

	// start server
	log.Printf("Server listening on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}

func getUserRouter(r *mux.Router, userHandler *handlers.UserHandler) {
	userRouter := r.PathPrefix("/api/users").Subrouter()
	userRouter.HandleFunc("/register", userHandler.RegisterUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/login", userHandler.LoginUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/refresh", userHandler.RefreshToken).Methods(http.MethodPost)

	protectedUserRouter := userRouter.PathPrefix("").Subrouter()
	protectedUserRouter.Use(middleware.Auth)
	protectedUserRouter.HandleFunc("/update", userHandler.UpdateUser).Methods(http.MethodPut)
	protectedUserRouter.HandleFunc("/update-profile-picture", userHandler.UpdateProfilePicture).Methods(http.MethodPut)
}
