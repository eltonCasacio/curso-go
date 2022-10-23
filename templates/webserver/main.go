package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Name         string
	CargaHoraria int
}

type Cursos []Curso

func toUpperCase(value string) string {
	return strings.ToUpper(value)
}
func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		t := template.Must(template.New("content.html").Funcs(template.FuncMap{"toUpper": toUpperCase}).ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{"Go", 120},
			{"Typescript", 170},
			{"Node", 140},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", mux)
}
