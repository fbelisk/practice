package main

import (
	"fmt"
	"log"
	"net/http"
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

func main() {
	go func() {
		for {
			log.Println("1")
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(n int) int {
	log.Println(n)
	return n
}
