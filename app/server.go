package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()
	//fmt.Println("l: ", l)
	for {
		connection, err := l.Accept()
		if err != nil {
			handleError(err)
		}

		var read []byte
		_, err = connection.Read(read)
		if err != nil {
			handleError(err)
		}

		connection.Write([]byte(("HTTP/1.1 200 OK\r\n\r\n")))

		connection.Close()
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
}
