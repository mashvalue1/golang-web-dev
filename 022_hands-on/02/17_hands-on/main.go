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

	switch {
	case mt == "GET" && uri == "/":
		serveIndex(conn)
	case mt == "GET" && uri == "/apply":
		serveApply(conn)
	case mt == "POST" && uri == "/apply":
		serveApplyPost(conn)
	default:
		serveDefault(conn)
	}
}

func serveIndex(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Index</title>
	</head>
	<body>
		<h1>Index</h1>
		<a href="/apply">apply</a>
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

func serveApply(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Apply</title>
	</head>
	<body>
		<h1>Apply</h1>
		<a href="/">index</a>
		<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
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

func serveApplyPost(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>ApplyPost</title>
	</head>
	<body>
		<h1>ApplyPost</h1>
		<a href="/">index</a><br>
		<a href="/apply">apply</a><br>
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

func serveDefault(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Default</title>
	</head>
	<body>
		<h1>Default</h1>
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
