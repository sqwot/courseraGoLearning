package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int `json:"user_id,string"`
	Username string
	Address  string `json:",omitempty"`
	Company  string `json:"-"`
}

var jsonStr = `[
	{"id":17, "username":"iivan","phone":0},
	{"id":"17", "address":"none", "company":"Mail.ru"}
]`

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************JSON***************")
	fmt.Println(string(colorReset), "")

	u := &User{
		ID:       42,
		Username: "sqwot",
		Address:  "test",
		Company:  "MARS Telecom",
	}

	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string: \n\t %s\n\n", result)

	data := []byte(jsonStr)
	var user1 interface{}
	json.Unmarshal(data, &user1)

	fmt.Printf("unpacked in empty interface:\n\t%#v\n\n", user1)
	user2 := map[string]interface{}{
		"id":       42,
		"username": "sqwotik",
	}
	var user2i interface{} = user2
	result2, err := json.Marshal(user2i)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string from map:\n\t%s\n\n", result2)
}
