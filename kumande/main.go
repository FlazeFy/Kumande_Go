package main

import (
	"kumande/packages/database"
	"kumande/routes"
)

func main() {
	database.Init()
	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":1323"))
}
