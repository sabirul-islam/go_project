package main

import (
	"fmt"
	"net"
	"os"
)


func main() {
	nl, err := net.Listen("tcp", ":8888")
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

	conn, err := nl.Accept()
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr)

	conn.Write([]byte("welcome to our coding bootcamp"))

	conn.Close()
	nl.Close()
}

/*
localhost
[::1]
127.0.0.1
192.168.88.172
*/

// telnet install: dism /online /Enable-Feature /FeatureName:TelnetClient