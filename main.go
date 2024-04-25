package main

import (
	"github.com/Osedisc/practiceGo/app"
	"github.com/Osedisc/practiceGo/infrastucture"
)

func init() {
	infrastucture.InitConfig()
}

func main() {
	app.Run()
}
