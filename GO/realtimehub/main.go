package main

import (
	"log"
	"os"

	"github.com/blue-jay/core/router"
	"github.com/blue-jay/core/server"
	"github.com/rs/cors"

	"github.com/chande/justasking/realtimehub/startup/boot"
	"github.com/chande/justasking/realtimehub/startup/env"
	"github.com/chande/justasking/realtimehub/startup/middleware"
)

func main() {

	// Set configuration file based on environment variable
	value := os.Getenv("JUSTASKING_ENV")
	configFile := "realtimehub/config.json"
	if value == "PROD" {
		configFile = "realtimehub/config.prod.json"
	}

	// Load the configuration file
	config, err := env.LoadConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	// Retrieve the middleware
	handler := middleware.SetUpMiddleware(router.Instance())

	// Register the services
	boot.RegisterServices()

	//set CORS and Methods Allowed
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:4200", "http://justasking.app"},
		AllowedHeaders: []string{
			"authorization",
			"cache-control",
			"content-type",
			"expires",
			"if-modified-since",
			"pragma",
			"strict-transport-security",
			"x-content-type-options",
			"x-frame-options",
			"x-xss-protection",
		},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})

	handler = c.Handler(handler)

	// Start the HTTP and HTTPS listeners
	server.Run(
		handler,       // HTTP handler
		handler,       // HTTPS handler
		config.Server, // Server settings
	)
}
