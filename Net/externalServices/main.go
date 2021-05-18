package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "getHandler: incoming request %#v\n,", r)
		fmt.Fprintf(w, "getHandler: r.URL %#v\n,", r.URL)
	})

	http.HandleFunc("/raw_body", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "postHandler: raw body %s\n", string(body))
	})
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

func runGet() {
	url := "http://localhost:8080/?param=123&param2=test"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("http.get body %#v\n\n\n", string(respBody))
}

func runGetFullReq() {
	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"User-Agent": {"coursera/golang"},
		},
	}
	req.URL, _ = url.Parse("http://localhost:8080/?id=42")
	req.URL.Query().Set("user", "sqwot")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error happend:", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("testGetFullReq resp %#v\n\n\n", string(respBody))

}

func runTransportAndPost() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	data := `{"id":42, "user":"sqwot"}`
	body := bytes.NewBufferString(data)

	url := "http://localhost:8080/raw_data"
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error happend:", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("runTransport %#v\n\n\n", string(respBody))
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************External Services***************")
	fmt.Println(string(colorReset), "")

	go startServer()

	time.Sleep(100 * time.Millisecond)

	runGet()
	runGetFullReq()
	runTransportAndPost()

}
