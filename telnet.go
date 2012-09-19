package main

import (
	"fmt"
	"log"
	"net"
)

func startSocket() {
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	log.Printf("About to start telnet://localhost:5000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		fmt.Println("Accepting a new connection")
		go doServeStuff(conn)
	}
}

func doServeStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading", err.Error())
			return
		}
		Publish(string(buf))
		conn.Write([]byte("ok\n"))
	}
}
