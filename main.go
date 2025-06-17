package main

import (
	"log"
	"net/http"

	"github.com/thilina01/kb-api-go/config"
	"github.com/thilina01/kb-api-go/routes"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	// Uncomment to seed (and comment it again after)
	// commands.SeedTags()

	routes.RegisterRoutes()

	handler := jsonMiddleware(http.DefaultServeMux)

	log.Println("🚀 KB API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
