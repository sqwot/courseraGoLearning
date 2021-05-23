package main

import (
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	ID       int
	Text     string
	Author   string
	Comments int
	Time     time.Time
}

func getPost(out chan []Post) {
	posts := []Post{}
	for i := 0; i < 10; i++ {
		post := Post{ID: 1, Text: "Text"}
		posts = append(posts, post)
	}
	out <- posts
}

func handleLeek(w http.ResponseWriter, r *http.Request) {
	res := make(chan []Post)
	go getPost(res)
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Gorutine trace***************")
	fmt.Println(string(colorReset), "")

	http.HandleFunc("/", handleLeek)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
