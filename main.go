package main

import (
	"fmt"
	"log"
	"time"

	"github.com/thiago-scherrer/metrobot/client"
	"github.com/thiago-scherrer/metrobot/lines"
)

func main() {
	for {
		body := client.Request()

		log.Println(lines.Blue(body))
		log.Println(lines.Green(body))
		log.Println(lines.Lilac(body))
		log.Println(lines.Red(body))
		log.Println(lines.Silver(body))
		log.Println(lines.Yellow(body))
		fmt.Println()

		time.Sleep(time.Duration(60) * time.Second)

	}
}
