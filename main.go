package main

import (
	"github.com/Gilberd-dev/task-5-pbi-btpns-gilberd-nicolas-siboro/database"
	"github.com/Gilberd-dev/task-5-pbi-btpns-gilberd-nicolas-siboro/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	r.Run(":8080")
}
