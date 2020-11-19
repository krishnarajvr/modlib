package main

import (
	"fmt"

	"github.com/krishnarajvr/modlib"
	"github.com/krishnarajvr/modlib/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", modlib.Config())
	fmt.Println(clientlib.Hello())
}
