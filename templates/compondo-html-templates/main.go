package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	t.Funcs(template.FuncMap{
		"toUpper": toUpper,
	})
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 45},
		{"Java", 60},
	})

	if err != nil {
		panic(err)
	}
}
