package main

import (
	"fmt"
	"ia04/agt"
)

func main() {
	ag := agt.NewRestServerAgent(":8080")
	ag.Start()
	fmt.Scanln()
}
