package main

import (
	"main.go/internel/app"
	"main.go/internel/router"
)

func main() {
	r := router.NewRouter()
	ap := app.NewApp(r)
	ap.Start()
}
