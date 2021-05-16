package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	name := conn.RemoteAddr().String()
	fmt.Printf("%+v connected\n\r", name)
	conn.Write([]byte("Hello, " + name + "\n\r"))
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Exit" {
			conn.Write([]byte("Bye\n\r"))
			fmt.Println(name, "disconnected")
			break
		} else if text != "" {
			fmt.Println(name, "enters", text)
			conn.Write([]byte("You enter: " + text + "\n\r"))
		}
	}
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************TCP***************")
	fmt.Println(string(colorReset), "")

	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
