package task

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func TaskSimulationFirst(seconds int) *http.Response {
	time.Sleep(time.Duration(seconds) * time.Second)

	res := &http.Response{Body: http.NoBody, Status: "100"}

	return res
}

func TaskSimulationSecond(seconds int) *http.Response {
	time.Sleep(time.Duration(seconds) * time.Second)

	res := &http.Response{Body: http.NoBody, Status: "100"}

	return res
}

func SyncExecApiCalls() {
	start := time.Now()

	t1 := TaskSimulationFirst(2)
	t2 := TaskSimulationSecond(3)

	total := convInt(t1.Status) + convInt(t2.Status)

	elapsed := time.Since(start)

	fmt.Println(total)
	fmt.Println(elapsed)
}

func AsyncExecApiCalls() {
	start := time.Now()

	c := make(chan *http.Response)

	go func() {
		t1 := TaskSimulationFirst(2)

		c <- t1
	}()

	go func() {
		t2 := TaskSimulationSecond(3)

		c <- t2
	}()

	t1, t2 := <-c, <-c
	total := convInt(t1.Status) + convInt(t2.Status)

	elapsed := time.Since(start)

	fmt.Println(total)
	fmt.Println(elapsed)
}

func convInt(value string) int {
	val, err := strconv.Atoi(value)

	if err != nil {
		return 0
	}

	return val
}
