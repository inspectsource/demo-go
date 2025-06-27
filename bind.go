package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func connect() {
	l, err := net.Listen("tcp", "0.0.0.0:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}

func x() {
	fmt.Println("Hello there")
}

func sshConfigure() {
	_ = ssh.InsecureIgnoreHostKey()
}
