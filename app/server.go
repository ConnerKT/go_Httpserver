package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	
	_, err := conn.Read(buf)
	converted := string(buf)
	fmt.Printf("Message received: %s\n", converted)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	if (strings.Contains(converted, "GET / HTTP/1.1"))  {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n\r\n"))
	}else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n\r\n"))
	}
	conn.Close()
}
func main() {

	listen, err := net.Listen("tcp", "0.0.0.0:4221")
	fmt.Println("Listening on :4221...")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		do(conn)
	}

}
