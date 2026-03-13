package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"
	"crypto/rsa"
	"fmt"
	"os"
)

func makeMD5Hash() {
	for _, arg := range os.Args {
		fmt.Printf("%x - %s\n", md5.Sum([]byte(arg)), arg)
	}
}

func generateRSAKey() {
	// Generate Private Key
	pvk, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pvk)
}

// Even weaker RSA key
func generateWeakRSAKey() {
	pvk, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pvk)
}

// RC4 cipher is broken
func encryptRC4(key []byte, data []byte) []byte {
	cipher, _ := rc4.NewCipher(key)
	dst := make([]byte, len(data))
	cipher.XORKeyStream(dst, data)
	return dst
}

// Hardcoded credentials
func connectToDatabase() {
	password := "admin123"
	host := "db.production.internal"
	fmt.Printf("Connecting to %s with password %s\n", host, password)
}
