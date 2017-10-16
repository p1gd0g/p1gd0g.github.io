package main

import (
	"fmt"
	"net"
)

func main() {

	var myPort string
	var hisPort string

	fmt.Println("your port")
	fmt.Scan(&myPort)

	fmt.Println("his port")
	fmt.Scan(&hisPort)

	l, _ := net.Listen("tcp", ":"+myPort)
	go func() {
		for {
			conn, _ := l.Accept()
			handle(conn)
		}
	}()

	for {
		message := make([]byte, 111)
		fmt.Scan(&message)
		conn, _ := net.Dial("tcp", ":"+hisPort)
		conn.Write(message)
	}

}

func handle(conn net.Conn) {
	message := make([]byte, 111)
	lengh, _ := conn.Read(message)

	fmt.Println(string(message[:lengh]))
}
