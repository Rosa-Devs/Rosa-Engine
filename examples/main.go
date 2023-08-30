package main

import (
	"log"

	"github.com/fatih/color"
	"github.com/libp2p/go-libp2p"
)

func main() {
	host, err := libp2p.New()
	if err != nil {
		panic("Fail create host: " + err.Error())
	}
	log.Println("Host created!")
	green := color.New(color.FgGreen).SprintFunc()
	log.Println("Host id:", green(host.ID().Pretty()))
	log.Println("Listening:", host.Addrs())

	//CREATE ROSA INSTANCE

}
