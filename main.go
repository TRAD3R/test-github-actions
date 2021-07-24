package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().Format("15:04:05"))
		time.Sleep(time.Second * 1)
	}
}
