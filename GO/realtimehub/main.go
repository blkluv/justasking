package main

import (
	"log"

	"github.com/blue-jay/core/router"
	"github.com/blue-jay/core/server"
	"github.com/rs/cors"

	"justasking/GO/realtimehub/startup/boot"
	"justasking/GO/realtimehub/startup/env"
	"justasking/GO/realtimehub/startup/middleware"
)

func main() {
	//enable when going to prod
	//appengine.Main()

	// Load the configuration file
	config, err := env.LoadConfig("config.json")
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
		AllowedOrigins:   []string{"http://localhost:4200", "http://justasking.io"},
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
