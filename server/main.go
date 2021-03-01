package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	nl, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}

	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(n)

	rqstr := string(bs)
	fmt.Println(rqstr)

	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>Helllo World</h1>
	</body>
	</html>`

	resp := "HTTP/1.1 200 ok\r\n" +
			"Content-Length: %d\r\n" +
			"Content-Type: text/html\r\n" +
			"\r\n%s"

	msg := fmt.Sprintf(resp, len(body), body)
	conn.Write([]byte(msg))
}