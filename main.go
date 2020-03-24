package main

import (
	"ScreenerDataServer/router"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		log.Print(e)
	}
	certfile := os.Getenv("certfile")
	keyfile := os.Getenv("keyfile")
	// create router and start listen on port 8000
	router := router.NewRouter()
	if certfile != "" && keyfile != "" {
		log.Fatal(http.ListenAndServeTLS(":6006", certfile, keyfile, setupGlobalMiddleware(router)))
	} else {
		log.Fatal(http.ListenAndServe(":6006", setupGlobalMiddleware(router)))
	}
}
