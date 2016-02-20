package lib

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	// "os/exec"
)

type OneCommand struct {
	Date string
}

func (self *OneCommand) Run(context *kingpin.ParseContext) (err error) {
	log.Printf("Context: %v", context)
	log.Printf("Date: %v", self.Date)
	return
}
