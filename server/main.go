package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
)

func main() {
	
	ln , err := net.Listen("tcp", ":8989")
	if err != nil {
		fmt.Println("error in listening to port", err);
		return ;
	}

	defer ln.Close()

	fmt.Println("server listening on 8989")

	for {

		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error in recieving the request from client ", err);
			return;
		}

		

		fmt.Println("Connection established to client at", conn.RemoteAddr())

		go handleConnection(conn)
	}
}


func handleConnection(conn net.Conn) {

	buf := make([]byte, 1024)
    msg := []byte{}

    for {
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("error in reading msg from client ", err)
			return ;
        }
        msg = append(msg, buf[:n]...)
        if bytes.Contains(buf[:n], []byte("\n")) {
            
        }else{
			continue
		}
		
		fmt.Println("client :", string(msg))
		msg = []byte{}

		fmt.Print("You : ")

		msg2 := bufio.NewReader(os.Stdin)
		msgBytes, err := msg2.ReadBytes('\n')

		conn.Write(msgBytes)
    }
}