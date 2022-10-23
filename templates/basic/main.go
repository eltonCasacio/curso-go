package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Name         string
	CargaHoraria int
}

func main() {
	curso := Curso{Name: "Go", CargaHoraria: 200}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Name}} - Carga Horaria: {{.CargaHoraria}} horas\n")

	f, err := os.Create("./templates/basic/template.txt")
	if err != nil {
		panic(err)
	}

	tmp.Execute(f, curso)
}
