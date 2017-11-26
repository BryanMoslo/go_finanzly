package main

import (
	"log"

	"github.com/BryanMoslo/go_finanzly/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
