//
// main.go
// dockswap
//
// Created by d-exclaimation on 00:00.
//

package main

import (
	"errors"
	"github.com/d-exclaimation/dockswap/cli"
	"github.com/d-exclaimation/dockswap/config"
	"log"
)

func main() {
	conf := config.Get()
	switch cli.Operation() {
	case cli.Initialize:
		conf.Init()
	case cli.Next:
		conf.Next()
	case cli.Error:
		log.Fatalln(errors.New("dockswap: Unrecognized argument"))
	case cli.Jump:
		args := cli.Args()
		if len(args) < 2 {
			log.Fatalln(errors.New("dockswap: Cannot jump undefined"))
		}
		conf.Jump(args[1])
	}
}
