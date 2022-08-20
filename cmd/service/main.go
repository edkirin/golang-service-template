package main

import (
	"golang-service-template/pkg/cfg"
	"golang-service-template/pkg/controller"
	"golang-service-template/pkg/db"
	"golang-service-template/pkg/logging"
)

func main() {
	cfg.Init()
	logging.Init()
	db.InitDB()
	controller.Serve()
}
