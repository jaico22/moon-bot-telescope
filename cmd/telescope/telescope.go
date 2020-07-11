package main

import (
	"time"

	"github.com/jaico22/moonbot-telescope/internal/telescope"
)

func main() {
	telescopeConfig := telescope.Config{
		SampleTime: time.Second * 5,
	}
	telescope.Setup(telescopeConfig)
	telescope.Run()
}
