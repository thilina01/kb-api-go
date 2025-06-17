package routes

import (
	"net/http"

	"github.com/thilina01/kb-api-go/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateTag(w, r)
		}
	})

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateArticle(w, r)
		}
	})
}
