package main

import (
	"fmt"
	"net/http"

	"cors-proxy/handler"
	"cors-proxy/logger"
	"cors-proxy/middleware"
)

func main() {
	var PORT uint16 = 8080

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Proxy)

	// Apply middleware in the correct order: logging -> CORS
	handlerWithMiddleware := middleware.Logging(middleware.CORS(mux))

	logger.Infof("Starting CORS proxy server on port %d", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), handlerWithMiddleware)
	if err != nil {
		logger.Errorf("Failed to start server: %v", err)
	}
}