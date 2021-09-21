package main

import (
	"final_project/configs"
	"final_project/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
}
