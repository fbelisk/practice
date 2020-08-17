package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
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
	t := make(map[int]int, 10)
	println(len(t))
	t[1] = 1
	println(len(t))
}

func Add(n int) int {
	log.Println(n)
	return n
}
