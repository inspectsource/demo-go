package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

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

// Unencrypted HTTP server
func startHTTPServer() {
	srv := &http.Server{
		Addr: ":80",
	}
	srv.ListenAndServe()
}

// Binding to all interfaces on privileged port
func bindAllInterfaces() {
	l, err := net.Listen("tcp", "0.0.0.0:22")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}

// SSH with password auth (weak)
func sshPasswordAuth() {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("password123"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	fmt.Println(config)
}
