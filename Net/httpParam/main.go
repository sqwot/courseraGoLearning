package main

import (
	"fmt"
	"net/http"
	"time"
)

var loginFromTmpl = []byte(`
	<html>
	<head>
		<title>lololo</title>
	</head>
	<body>
		<form action="/" method="post">
			login: <input name="login" type="text"/>
			password: <input name="password" type="password"/>
			<input type="submit" value="go">
		</form>
	</body>
	</html>
`)

func header(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("RequestID", "sd1wf243rv3ba")

	fmt.Fprintln(w, "your browser is", r.UserAgent())
	fmt.Fprintln(w, "you accept", r.Header.Get("Accept"))
}
func loginPage(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   "sqwot",
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
func logoutPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	http.Redirect(w, r, "/", http.StatusFound)
}

func handler(w http.ResponseWriter, r *http.Request) {
	myParam := r.URL.Query().Get("param")
	if myParam != "" {
		fmt.Fprintln(w, "`myparam` is", myParam)
	}

	key := r.FormValue("key")
	if key != "" {
		fmt.Fprintln(w, "`key` is", key)
	}
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	w.Write(loginFromTmpl)
	// 	return
	// }
	// //r.ParseForm()
	// //inputLogin := r.Form["login"][0]

	// inputLogin := r.FormValue("login")
	// fmt.Fprintln(w, "you enter: ", inputLogin)

	session, err := r.Cookie("session_id")
	loggedIn := (err != http.ErrNoCookie)
	if loggedIn {
		fmt.Fprintln(w, `<a href="/logout">logout</a>`)
		fmt.Fprintln(w, "Wlcome, "+session.Value)
	} else {
		fmt.Fprintln(w, `<a href=/login>login</a>`)
		fmt.Fprintln(w, "you need to login")
	}
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Парамерты запроса***************")
	fmt.Println(string(colorReset), "")

	http.HandleFunc("/handler", handler)
	http.HandleFunc("/header", header)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/logout", logoutPage)
	http.HandleFunc("/", mainPage)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)

}
