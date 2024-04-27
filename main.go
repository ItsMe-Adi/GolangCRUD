package main

import (
	"playground/api"
	"playground/database"
	"playground/util"
)

func main() {
	util.LoadConfig(".")
	database.InitDB()
	api.StartServer()
}
