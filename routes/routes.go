package routes

import (
	"net/http"

	"github.com/thilina01/kb-api-go/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.CreateTag(w, r)
		case http.MethodGet:
			controllers.ListTags(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			controllers.CreateArticle(w, r)
		case http.MethodGet:
			controllers.ListArticles(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			controllers.UpdateArticle(w, r)
		case http.MethodDelete:
			controllers.DeleteArticle(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
