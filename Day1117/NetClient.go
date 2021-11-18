package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	connnect, _ := net.Dial("tcp","localhost:8888")
	fmt.Println(connnect)
	log.Println(connnect)
	defer connnect.Close()
	io.Copy(os.Stdout,connnect)
}