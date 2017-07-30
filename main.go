package main

import (
	"github.com/yundream/contextlogging/app"
)

func main() {
	LoggingApp := app.New()
	LoggingApp.Run(":8000")
}
