package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		handleError(err)
	}
	defer l.Close()
	//fmt.Println("l: ", l)
	for {
		connection, err := l.Accept()
		if err != nil {
			handleError(err)
		}
		// process request
		request := make([]byte, 1024)
		_, err = connection.Read(request)
		if err != nil {
			handleError(err)
		}
		requestData := string(request)
		fmt.Printf("Data: \n%s", requestData)
		splitData := strings.Split(requestData, "\r\n")
		for i, datum := range splitData {
			fmt.Printf("i=%v -- Data: %v----\n", i, datum)
		}

		startLine := parseStartLine(splitData[0])

		if startLine.Path == "/" {
			connection.Write([]byte(("HTTP/1.1 200 OK\r\n\r\n")))
		} else if strings.HasPrefix(startLine.Path, "/echo") {
			res := echoResponse(startLine.Path)
			connection.Write([]byte((res)))
		} else {
			connection.Write([]byte(("HTTP/1.1 404 NOT FOUND\r\n\r\n")))

		}

		connection.Close()
	}
}

func echoResponse(path string) string {
	content := strings.TrimPrefix(path, "/echo/")
	response := make([]string, 5)
	response[0] = "HTTP/1.1 200 OK"
	response[1] = "Content-Type: text/plain"
	response[2] = fmt.Sprintf("Content-Length: %d", len(content))
	response[3] = CONTENT_SEPARATOR
	response[4] = content
	return strings.Join(response, LINE_SEPARATOR)

}

func parseStartLine(line string) StartLine {
	items := strings.Split(line, " ")
	if len(items) != 3 {
		log.Fatal("Expect 'HTTP_METHOD<space>PATH<space>HTTP_VERSION'")
	}
	return StartLine{
		HttpMethod:  items[0],
		Path:        items[1],
		HttpVersion: items[2],
	}

}

type StartLine struct {
	HttpMethod  string
	Path        string
	HttpVersion string
}

//goland:noinspection GoSnakeCaseUsage
const LINE_SEPARATOR = "\r\n"

//goland:noinspection GoSnakeCaseUsage
const CONTENT_SEPARATOR = ""

func handleError(err error) {
	if err != nil {
		fmt.Println("error occurred: ", err.Error())
		os.Exit(1)
	}
}
