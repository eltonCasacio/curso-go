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
	f, err := os.Create("./templates/must/template.txt")
	if err != nil {
		panic(err)
	}

	tmp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Name}} - Carga Horaria: {{.CargaHoraria}} horas\n"))
	err = tmp.Execute(f, curso)
	if err != nil {
		panic(err)
	}
}
