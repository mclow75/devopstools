package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed migrations/*.sql
var migrations embed.FS

//go:embed templates/*.html
var templateFS embed.FS

func main() {
	xs, err := migrations.ReadDir("migrations")
	if err != nil {
		panic("migrations folder or file not includes")
	}
	for _, x := range xs {
		fmt.Println("Migrations:" + x.Name())
	}

	var templates = template.Must(template.ParseFS(templateFS, "templates/*.html"))

	router := http.NewServeMux()

	router.HandleFunc("/", func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		templates.ExecuteTemplate(
			w,
			"root.html",
			map[string]string{
				"Title": "Foo",
			})
	})

	http.ListenAndServe(":8800", router)
}
