package main

import (
	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/middleware"
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

	// initialize services

	// initialize handlers

	// setup routes
	r := mux.NewRouter()
	r.Use(middleware.CORS)

	// start server
	log.Printf("Server listening on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
