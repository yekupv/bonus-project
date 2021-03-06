package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter your name: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(conn, text+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}

		fmt.Print("Hello, " + message)
	}
}
