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

	connection, err := l.Accept()
	if err != nil {
		handleError(err)
	}
	var read []byte
	connection.Read(read)

	connection.Write([]byte(("HTTP/1.1 200 OK\r\n\r\n")))

	connection.Close()

}

func handleError(err error) {
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
}
