package main

import (
	//"fmt"
	"net"
	//	"os"
	//"time"
)

const (
	policyString = `<?xml version="1.0"?>
<cross-domain-policy>
<allow-access-from domain="*" to-ports="*" />
</cross-domain-policy>
`
)

func main() {
	port := ":843"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkError(err)

	println("Lanching a server...")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.Write([]byte(policyString))
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
