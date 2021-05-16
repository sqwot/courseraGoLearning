package main

import (
	"fmt"
	"net/http"
	"time"
)

const port = ":8080"

type Handler struct {
	Name string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Name:", h.Name, "URL:", r.URL.String())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "handler one")
	w.Write([]byte("!!!"))
}
func handler2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("handler two"))
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Simple WebServer***************")
	fmt.Println(string(colorReset), "")

	testHandler := &Handler{Name: "testy test"}
	http.Handle("/test/", testHandler)

	//rootHandler := &Handler{Name: "rooty root"}
	//http.Handle("/", rootHandler)

	//http.HandleFunc("/", handler)
	http.HandleFunc("/again", handler2)
	http.HandleFunc("/pages", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Single page:", r.URL.String())
	})
	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Multiple page:", r.URL.String())
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("starting server at %s", port)
	//http.ListenAndServe(port, nil)
	server.ListenAndServe()
}
