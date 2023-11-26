package main

import (
	"github.com/BaptisteBuvron/go-voting/agt"
)

func main() {
	ag := agt.NewRestServerAgent(":8080")
	ag.Run()
}
