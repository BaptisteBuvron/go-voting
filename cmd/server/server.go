package main

import (
	"ia04/agt"
)

func main() {
	ag := agt.NewRestServerAgent(":8080")
	ag.Run()
}
