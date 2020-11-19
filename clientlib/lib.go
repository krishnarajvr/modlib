package main

import (
	"fmt"

	"github.com/eliben/modlib"
	"github.com/eliben/modlib/clientlib"
)

func main() {
	fmt.Println(modlib.Config())
	fmt.Println(clientlib.Hello())
}
