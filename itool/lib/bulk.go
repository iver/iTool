package lib

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	// "os/exec"
)

type BulkCommand struct {
	InitDate string
	EndDate  string
}

func (self *BulkCommand) Run(context *kingpin.ParseContext) (err error) {
	log.Printf("Context: %v", context)
	log.Printf("InitDate: %v and EndDate: %v", self.InitDate, self.EndDate)
	return
}
