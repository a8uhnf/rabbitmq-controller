package main

import (
	"log"
	
	"github.com/a8uhnf/rabbitmq-controller/cmd"
)

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		log.Println(err)
	}
}
