package main

import (
	"log"

	"github.com/bperezgo/go-template/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Panic(err)
	}
}
