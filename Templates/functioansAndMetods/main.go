package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func (u *User) PrintActive() string {
	if !u.Active {
		return ""
	}
	return "method says user " + u.Name + " active"
}

func IsUserOdd(u *User) bool {
	return u.ID%2 != 0
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Metods and Functions on Template***************")
	fmt.Println(string(colorReset), "")

	tmpl, err := template.New("").ParseFiles("method.html")
	if err != nil {
		panic(err)
	}

	tmplFuncs := template.FuncMap{
		"OddUser": IsUserOdd,
	}
	tmpl2, err := template.New("").Funcs(tmplFuncs).ParseFiles("func.html")
	if err != nil {
		panic(err)
	}

	users := []User{
		User{1, "Vasiliy", true},
		User{2, "Ivan", false},
		User{3, "Dmitry", true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "method.html", struct{ Users []User }{users})
		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("/func", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl2.ExecuteTemplate(w, "func.html", struct{ Users []User }{users})
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)

}
