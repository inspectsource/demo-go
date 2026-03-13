package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	productPath string = "/tmp/products.json"
	sellerPath  string = "/tmp/sellers.json"
)

func main() {
	s := Seller{
		Name: "SomeSeller",
		Address: Address{
			City: "Gotham",
		},
	}
	var i GenericInterface
	i = &s

	givenCity := "Gotham"

	switch i.(type) {
	case interface{}:
		fmt.Println("What to do")
	case Product:
		fmt.Println(i.(Product).DeliversTo(givenCity))
	case Seller:
		fmt.Println(i.(Seller).DeliversTo(givenCity))
	}

	// Command injection via user input
	userInput := os.Args[1]
	runCmd("bash", []string{"-c", userInput}, nil, ".")

	// Useless string comparison
	name := "hello"
	if strings.Compare(name, "world") == 0 {
		fmt.Println("equal")
	}

	// Dead code: variable assigned but never used
	unusedVar := 42
	_ = unusedVar
}

func appendData() {
	var s []string
	s = append(s, productPath)
	s = append(s, sellerPath)
}

// Receiver name inconsistency and stutter
type DataData struct {
	Value string
}

func (d DataData) GetValue() string {
	return d.Value
}

func (self DataData) SetValue(v string) string {
	self.Value = v
	return self.Value
}
