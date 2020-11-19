package main

import (
	"fmt"

	"github.com/krishnarajvr/modlib"
	"github.com/krishnarajvr/modlib/internal/auth"
	"github.com/krishnarajvr/modlib/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", modlib.Config())
	fmt.Println("Auth:", auth.GetAuth())
	fmt.Println(serverlib.Hello())
}
