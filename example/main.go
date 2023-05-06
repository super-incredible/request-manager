package main

import (
	"fmt"

	"github.com/super-incredible/request-manager/example/task"
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

func main() {
	secondAsync := 3
	fmt.Printf("Waiting request for %ss... ", fmt.Sprint(secondAsync))

	task.AsyncExecApiCalls()

	fmt.Println("--------------")

	secondSync := 5
	fmt.Printf("Waiting request for %ss... ", fmt.Sprint(secondSync))

	task.SyncExecApiCalls()

	fmt.Println("--------------")

	fmt.Println("Done")
}
