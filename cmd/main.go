package main

import (
	"log"

	"github.com/bperezgo/testing-ai/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Panic(err)
	}
}
