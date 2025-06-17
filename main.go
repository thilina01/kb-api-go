package main

import (
	"log"
	"net/http"

	"github.com/thilina01/kb-api-go/config"
	"github.com/thilina01/kb-api-go/routes"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	routes.RegisterRoutes()

	log.Println("🚀 KB API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
