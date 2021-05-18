package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var uploadFormTmpl = []byte(`
<html>
	<body>
	<form action="/upload" method="post" enctype="multipart/form-data">
		Image: <input type="file" name="my_file">
		<input type="submit" value="Upload">
	</form>
	</body>
</html>		
`)

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write(uploadFormTmpl)
}
func uploadPage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1025)
	file, handler, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename)
	fmt.Fprintf(w, "handler.Header %#v\n", handler.Header)
	fmt.Fprintf(w, "handler.Size %d\n", handler.Size)

	hasher := md5.New()
	io.Copy(hasher, file)

	fmt.Fprintf(w, "md5 %x\n", hasher.Sum(nil))
}

type Params struct {
	ID   int
	User string
}

/*
curl -v -X POST -H "Content-Type: application/json" -d '{"id":2, "user":"sqwot"}' http://localhost:8080/raw_body
*/

func uploadRawBody(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	fmt.Println(string(body))
	p := &Params{}
	err = json.Unmarshal(body, p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Content-Type %#v\n", r.Header.Get("Content-Type"))
	fmt.Fprintf(w, "params %#v\n", p)
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Loading files from form***************")
	fmt.Println(string(colorReset), "")

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/raw_body", uploadRawBody)
	http.HandleFunc("/upload", uploadPage)

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
