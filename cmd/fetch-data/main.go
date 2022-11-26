package main

import (
	"fmt"

	"github.com/ristoavila/nilan"
)

func main() {
	c := nilan.Controller{Config: nilan.Config{NilanAddress: "192.168.0.42:502"}}
	errors, _ := c.FetchErrors()
	fmt.Println(errors)
}
