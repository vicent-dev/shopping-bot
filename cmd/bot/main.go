package main

import (
	"log"
	"shopping-bot/app"
)

func main() {
	s := app.NewServer()

	if err := s.Run(); s != nil {
		log.Fatalln(err)
	}
}
