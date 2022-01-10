package main

import (
	"brian.com/upstart/internal/config"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello Go")
	config, err := config.GetAppConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(config.DBTable)
}
