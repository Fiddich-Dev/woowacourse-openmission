package main

import (
	"golang-server/internal/routers"
)

func main() {
	app := routers.SetupRouter()
	app.Run(":8080")
}
