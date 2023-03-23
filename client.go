package main

import (
	"net"
	"log"
	"fmt"
	// "netpackets4go/utils"
)

const (
	HOST = "localhost"
	PORT = "8023"
	TYPE = "tcp"
)

func main() {
	var str string
	var err error
	if str, err = client_func(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}

func client_func() (string, error) {
	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		return "", err
	}

	_, err = conn.Write([]byte("recieved by client; terminating client network connection."))
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 2048)

	var ret_value string = ""
	// mlen, _ := conn.Read(buffer)
	mlen, _ := conn.Read(buffer)
	// fmt.Println(mlen)
	// fmt.Println(buffer)
	ret_value += string(buffer[0:mlen])

	return ret_value, err
}
