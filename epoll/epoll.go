package epoll

import (
	"fmt"
	"net"
)

func CreateEpoll() {
	a := Poll{}
	fmt.Println(a)
	l , err := net.Listen()
	l.Accept()
}

func AddRead() {

}

