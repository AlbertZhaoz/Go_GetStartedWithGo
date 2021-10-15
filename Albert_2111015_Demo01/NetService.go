package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp","localhost:8888")
	if(err!=nil){
		log.Fatal(err)
	}

	for{
		connect,err := listener.Accept()
		if(err!=nil){
			log.Fatal(err)
		}

		fmt.Println("A client connected.")
		go handle(connect) //启了一个协程
	}
}

func handle(connect net.Conn)  {
	defer connect.Close()

	for i := 0; i < 10; i++ {
		_,err := io.WriteString(connect,"Elegant.\n")
		if err!=nil{
			log.Fatal(err)
		}
		time.Sleep(1*time.Second)
	}
}