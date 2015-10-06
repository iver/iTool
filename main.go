package main

import (
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("iTool", "A command-line tool for iTunes Connect Reports")

	vendorId   = app.Flag("vendor_id", "Apple vendor id.").Short('V').Required().String()
	properties = app.Flag("properties_filename", "Name of the properties file, iTool use bin/autoingestion.properties by default.").Short('f').Default("bin/autoingestion.properties").String()
	reportType = app.Flag("report_type", "Sales (Default) or Newsstand.").Short('t').Default("Sales").String()
	dateType   = app.Flag("date_type", "Daily (Default), Weekly, Monthly, Yearly").Short('p').Default("Daily").String()

	bulk     = app.Command("bulk", "Fetch iTunes Connect Reports on range of dates")
	initDate = bulk.Arg("init_date", "Initial report date on YYYYMMDD format").Required().String()
	endDate  = bulk.Arg("end_date", "End report date on YYYYMMDD format").Required().String()

	byOne = app.Command("one", "Fetch one iTunes Connect Report from specific date")
	date  = byOne.Arg("date", "Date on YYYYMMDD format").String()
)

func main() {
	app.Version("iTool. Version: 0.0.1")
	kingpin.MustParse(app.Parse(os.Args[1:]))
	log.Printf("salida | vendorId: %v, properties: %v, reportType: %v, dateType: %v", *vendorId, *properties, *reportType, *dateType)
}
