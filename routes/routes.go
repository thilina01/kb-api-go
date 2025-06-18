package routes

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/thilina01/kb-api-go/controllers"
	_ "github.com/thilina01/kb-api-go/docs"
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
		case http.MethodGet:
			controllers.GetArticleByID(w, r)
		case http.MethodPut:
			controllers.UpdateArticle(w, r)
		case http.MethodDelete:
			controllers.DeleteArticle(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/articles/search", controllers.SearchArticles)

	http.HandleFunc("/ping", controllers.Ping)

	http.Handle("/swagger/", httpSwagger.WrapHandler)
	// ...
}
