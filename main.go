package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"struct/tcpClient"
	"struct/tcpServer"
	"struct/utils"
	"time"
)

func deadloop() {
	for {
		_ = 1
	}
}

func worker() {
	for {
		fmt.Println("worker is running")
		time.Sleep(time.Second * 1)
	}
}

type Server struct {
	Name        string `json:"name,omitempty"`
	Port        int    `json:"port"`
	EnableLogs  bool   `json:"enable_logs"`
	BaseDomain  string `json:"base_domain"`
	Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

func (receiver *Server) name()  {
	receiver.Name = "ssss"
}

func main() {
	w := &util.WaitGroupWrapper{}
	w.Wrap(tcpServer.Run)
	time.Sleep(2 * time.Second)
	w.Wrap(tcpClient.Run)
	w.Wait()
}

func Add(n int) int {
	log.Println(n)
	return n
}
