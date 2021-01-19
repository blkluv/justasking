package main

import (
	"justasking/GO/api/controller"
	"log"

	"github.com/blue-jay/core/router"
	"github.com/blue-jay/core/server"
	"github.com/rs/cors"

	"justasking/GO/api/startup/middleware"
	"justasking/GO/core/startup/boot"
	"justasking/GO/core/startup/env"
)

func main() {
	//enable when going to prod
	//appengine.Main()

	// Load the configuration file
	config, err := env.LoadConfig("config.json")
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
			"http://justasking.io",
			"https://justasking.io",
			"http://www.justasking.io",
			"https://www.justasking.io",
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
