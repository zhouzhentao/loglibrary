package main

import (
	"fmt"
	"time"

	"loglibrary/src/log.v1"
)

func main() {
	log.InitLoger("console")

	for {
		time.Sleep(time.Second)
		fmt.Println("sleep 1 second")
		go log.Warn("i am a Warn log\n")
	}
}
