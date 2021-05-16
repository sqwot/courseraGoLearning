package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

var xmlData = []byte(`asdasdasd`)

type User struct {
	ID      int    `xml:"id,attr"`
	Login   string `xml:"login"`
	Name    string `xml:"name"`
	Browser string `xml:"browser"`
}

type Users struct {
	Version string `xml:"version, attr"`
	List    []User `xml:"user"`
}

func CountStruct() {
	logins := make([]string, 0)
	v := new(Users)
	err := xml.Unmarshal(xmlData, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	for _, u := range v.List {
		logins = append(logins, u.Login)
	}
}
func CountDecoder() {
	input := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(input)
	logins := make([]string, 0)
	var login string
	for {
		tok, tokenErr := decoder.Token()
		if tokenErr != nil && tokenErr != io.EOF {
			fmt.Println("error happend", tokenErr)
			break
		} else if tokenErr == io.EOF {
			break
		}
		if tok == nil {
			fmt.Println("t is nil break")
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "login" {
				if err := decoder.DecodeElement(&login, &tok); err != nil {
					fmt.Println("error happend", err)
				}
				logins = append(logins, login)
			}
		}
	}
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************XML***************")
	fmt.Println(string(colorReset), "")

	CountStruct()
	CountDecoder()

}
