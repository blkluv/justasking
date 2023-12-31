package main

import (
	"log"
	"os"

	"github.com/chande/justasking/api/controller"

	"github.com/blue-jay/core/router"
	"github.com/blue-jay/core/server"
	"github.com/rs/cors"

	"github.com/chande/justasking/api/startup/middleware"
	"github.com/chande/justasking/core/startup/boot"
	"github.com/chande/justasking/core/startup/env"
)

func main() {

	// Set configuration file based on environment variable
	value := os.Getenv("JUSTASKING_ENV")
	configFile := "api/config.json"
	if value == "PROD" {
		configFile = "api/config.prod.json"
	}

	// Load the configuration file
	config, err := env.LoadConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	// Register the services
	boot.RegisterServices(config)

	// Load the controller routes
	controller.LoadRoutes()

	// Retrieve the middleware
	handler := middleware.SetUpMiddleware(router.Instance())

	//set CORS and Methods Allowed
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins: []string{
			"http://localhost:4200",
			"http://justasking.app",
			"https://justasking.app",
			"http://www.justasking.app",
			"https://www.justasking.app",
		},
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
			"origin",
			"x-requested-with",
			"accept",
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
