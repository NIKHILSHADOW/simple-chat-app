package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)


func main() {

	serverAddr := "localhost:8989"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	
	
	for {

		fmt.Print("you : ")

		msg := bufio.NewReader(os.Stdin)

		msgBytes, err := msg.ReadBytes('\n')

		d, err := conn.Write(msgBytes)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(d)

		buf := make([]byte, 1024)
		msg2 := []byte{}
		
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		msg2 = append(msg2, buf[:n]...)
		
		if bytes.Contains(msg2, []byte("\n")) {

		}else {
			continue
		}

		fmt.Println("server : ", string(msg2))
		msg2 = []byte{}

	}
}