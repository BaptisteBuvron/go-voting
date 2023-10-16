package main

import (
	"fmt"
	"tp3/agt"
)

func main() {
	ag := agt.NewRestServerAgent(":8080")
	ag.Start()
	fmt.Scanln()
}
