package tcpServer

import (
	"fmt"
	"net"
)

func Server(){
	listener, err := net.Listen("tcp", ":10101")
	if err != nil {
		fmt.Println("HOLY SHIT IT DOES NOT RUN SUCCESS")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept conn err "+ err.Error())
			continue
		}
		conn.Read()
	}
}

