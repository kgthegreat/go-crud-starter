package routes

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/kgthegreat/go-crud-starter/app"
	"github.com/kgthegreat/go-crud-starter/controllers"
	"github.com/kgthegreat/go-crud-starter/repositories"
)

func NewRouter(a *app.App) *chi.Mux {
	//	r := mux.NewRouter()
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Repositories
	topicRepository := repositories.NewTopicRepository(a.Database)

	// Controllers
	topicController := controllers.NewTopicController(a, topicRepository)

	distDir := a.Config.StaticDir
	FileServer(r, "/static", http.Dir(distDir))

	r.Get("/", topicController.GetAll)

	r.Route("/topics", func(r chi.Router) {

		r.Post("/", topicController.Create)                // POST /topics
		r.Get("/{topicId}", topicController.GetById)       // GET /topics/:id
		r.Delete("/{topicId}", topicController.DeleteById) // DELETE /topics/:id
	})

	// Topics
	// r.HandleFunc("/", topicController.GetAll).Methods(http.MethodGet)
	// r.HandleFunc("/topics/{id:[0-9]+}", topicController.GetById).Methods(http.MethodGet)
	// r.HandleFunc("/topics", topicController.Create).Methods(http.MethodPost)

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
