package lib

import "gopkg.in/alecthomas/kingpin.v2"

const (
	desc          = "A command-line tool for iTunes Connect Reports"
	vendorMsg     = "Apple vendor id"
	propertiesMsg = "Name of the properties file, iTool use bin/autoingestion.properties by default."
	reportTypeMsg = "Sales (Default) or Newsstand."
	dateTypeMsg   = "Daily (Default), Weekly, Monthly, Yearly"

	bulkMsg     = "Fetch iTunes Connect Reports on range of dates"
	initDateMsg = "Initial report date on YYYYMMDD format"
	endDateMsg  = "End report date on YYYYMMDD format"

	oneMsg  = "Fetch one iTunes Connect Report from specific date"
	dateMsg = "Date on YYYYMMDD format"
	Version = "iTool Version: 0.0.1"
)

type App struct {
	app            *kingpin.Application
	Log            *Logger
	Config         *Config
	VendorId       string
	PropertiesFile string
	ReportType     string
	DateType       string
}

func NewApp() (application *App) {
	config, _ := NewConfig()
	application = &App{
		app:    kingpin.New("iTool", desc),
		Config: config,
		Log:    NewLogger(config),
	}
	application.app.Version(Version)
	application.required()
	return
}

func (self *App) Parse(args []string) {
	kingpin.MustParse(self.app.Parse(args))
}

func (self *App) required() {
	self.app.Flag("vendor_id", vendorMsg).Short('V').Required().StringVar(&self.VendorId)
	self.app.Flag("properties_filename", propertiesMsg).Short('f').Default("vendor/autoingestion.properties").StringVar(&self.PropertiesFile)
	self.app.Flag("report_type", reportTypeMsg).Short('t').Default("Sales").StringVar(&self.ReportType)
	self.app.Flag("date_type", dateTypeMsg).Short('p').Default("Daily").StringVar(&self.DateType)
}

func (self *App) Bulk(cmd BulkCommand) {
	bulk := self.app.Command("bulk", bulkMsg).Action(cmd.Run)
	bulk.Arg("init_date", initDateMsg).Required().StringVar(&cmd.InitDate)
	bulk.Arg("end_date", endDateMsg).Required().StringVar(&cmd.EndDate)
}

func (self *App) One(cmd OneCommand) {
	one := self.app.Command("one", oneMsg).Action(cmd.Run)
	one.Arg("date", dateMsg).StringVar(&cmd.Date)
}
