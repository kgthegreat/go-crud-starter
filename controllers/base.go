package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kgthegreat/go-crud-starter/models"
)

type ResponseData struct {
	Topics []*models.Topic
}

type AppError struct {
	Success bool
	Message string
	Status  int
	e       error
}

func NewAppError(e *AppError, w http.ResponseWriter) {
	if e.e != nil {
		log.Println(e.e)
	}

	http.Error(w, e.e.Error(), http.StatusInternalServerError)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	templates := template.Must(template.New("").Funcs(template.FuncMap{
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
	}), nil)

	filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			templates.ParseFiles(path)
		}

		return nil
	})

	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func retrieveFormData(r *http.Request, w http.ResponseWriter, field string) string {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	topic := r.Form.Get(field)
	return topic

}
