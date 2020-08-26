package tcpServer

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	util "struct/utils"
)

func Run(){
	listener, err := net.Listen("tcp", ":10101")
	if err != nil {
		fmt.Println("HOLY SHIT IT DOES NOT RUN SUCCESS")
	}
	w := &util.WaitGroupWrapper{}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept conn err "+ err.Error())
			continue
		}
		w.Wrap(func() {
			for {
				packLen := make([]byte, 4)
				_, err = io.ReadFull(conn, packLen)
				if err != nil {
					fmt.Println("read length error "+ err.Error())
					break
				}
				length := binary.BigEndian.Uint32(packLen)
				bodyByte := make([]byte, length)
				_, err = io.ReadFull(conn, bodyByte)
				if err != nil {
					fmt.Println("read body error "+ err.Error())
					break
				}
				fmt.Println("msg " + string(bodyByte))
			}
		})
	}
	w.Wait()
}

