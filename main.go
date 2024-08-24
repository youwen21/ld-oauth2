package main

import (
	"fmt"
	"gofly/conf"
	"gofly/conf/flag_vars"
	"gofly/server"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func handleSignals() {
	//pid := os.Getpid()
	//_ = os.WriteFile("pid", []byte(strconv.Itoa(pid)), 777)
	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM) //NOTE:: syscall.SIGKILL we cannot catch kill -9 as its force kill signal.
	<-terminateSignals
	//_ = os.Remove("pid")
}

func main() {
	fmt.Println("starting")

	listener, err := net.Listen("tcp", conf.GetAddress())
	if err != nil {
		log.Fatal("init listener failure:", err)
	}
	log.Println("listen at:", conf.GetAddress())
	log.Println("local open:", fmt.Sprintf("http://127.0.0.1:%v", flag_vars.Port()))

	go func() {
		err = http.Serve(listener, server.Mux)
		if err != nil {
			log.Fatal("http.Serve failure:", err)
		}
	}()

	handleSignals()
}
