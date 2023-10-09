package main

import (
	"fmt"
	"tp3/agt"
)

func main() {
	ag := agt.NewRestClientAgent("id1", "http://localhost:8080")
	ag.Start()
	fmt.Scanln()
}
