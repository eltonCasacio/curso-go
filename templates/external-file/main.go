package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Name         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	f, err := os.Create("./templates/external-file/html-gerado.html")
	tmp := template.Must(template.New("template.html").ParseFiles("./templates/external-file/template.html"))
	err = tmp.Execute(f, Cursos{
		{"Go", 120},
		{"Typescript", 170},
		{"Node", 40},
	})
	if err != nil {
		panic(err)
	}
}
