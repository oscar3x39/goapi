package main

import (
	"goapi/models"
	"goapi/routers"
)

func main() {
	models.ConnectDataBase()
	r := routers.SetupRouter()
	r.Run()
}