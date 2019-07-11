package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var i int
	var mt, uri string

	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			// request header
			rq := strings.Fields(line)
			mt = rq[0]
			uri = rq[1]
			fmt.Println(mt)
			fmt.Println(uri)
		}
		if line == "" {
			break
		}
		i++
	}

	// body := "<h1>"
	// body += "Response Body"
	// body += "\n"
	// body += mt
	// body += "\n"
	// body += uri
	// body += "</h1>"
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
		</body>
		</html>
	`

	// response header
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	// response body
	io.WriteString(conn, body)

}
