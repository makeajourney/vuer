package main

import (
	"log"

	"github.com/makeajourney/vuer/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
