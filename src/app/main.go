package main

import (
	"time"

	"loglibrary/src/log.v1"
)

func main() {
	log.InitLoger("file")
	go log.Fatal("i am a Fatal log\n")
	for {
		time.Sleep(time.Second)
	}

}
