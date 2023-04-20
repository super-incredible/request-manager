package main

import (
	"fmt"
	"time"
)

type RequestDependency struct {
	Name      string   `json:"name"`
	Variables []string `json:"variables"`
}

type Request struct {
	Name         string              `json:"name"`
	Dependencies []RequestDependency `json:"dependencies"`
}

type Requests []Request

func asyncSimulation(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	seconds := 3
	fmt.Printf("Waiting request for %ss... ", fmt.Sprint(seconds))

	asyncSimulation(seconds)

	fmt.Println("Done")
}
