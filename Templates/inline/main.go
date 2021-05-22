package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Templates inline***************")
	fmt.Println(string(colorReset), "")

	http.HandleFunc("/", handle)

	tmpl := template.Must(template.ParseFiles("users.html"))

	users := []User{
		User{1, "Vasiliy", true},
		User{2, "<i>Ivan</i>", false},
		User{3, "Dmitry", true},
	}

	http.HandleFunc("/html/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Users []User }{
			users,
		})
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

type tplParams struct {
	URL     string
	Browser string
}

type User struct {
	ID     int
	Name   string
	Active bool
}

const EXAMPLE = `
Browser {{.Browser}}

you at {{.URL}}
`

func handle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New(`example`)
	tmpl, _ = tmpl.Parse(EXAMPLE)

	params := tplParams{
		URL:     r.URL.String(),
		Browser: r.UserAgent(),
	}

	tmpl.Execute(w, params)
}
