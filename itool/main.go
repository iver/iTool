package main

import (
	"os"

	"itool/lib"
)

func main() {
	var app = lib.NewApp()
	app.One(lib.OneCommand{})
	app.Bulk(lib.BulkCommand{})
	app.Parse(os.Args[1:])
	app.Log.Printf(" --- Finalize: %v --- ", lib.Version)
}
