package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thilina01/kb-api-go/config"
)

func main() {
	err := config.ConnectDB()
	if err != nil {
		log.Fatal("❌ MongoDB Connection Failed:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Knowledge Base API is running!")
	})

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
