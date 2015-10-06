package main

import (
	"os"

	"github.com/ivan-iver/itool/lib"
)

func main() {
	var app = lib.NewApp()
	app.Parse(os.Args[1:])
	app.Log.Printf(" --- Finalize: %v --- ", lib.Version)
}
