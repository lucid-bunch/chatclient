package main

import (
	"chatclient/app"
)

func main() {
	app := app.NewApp("localhost:3000")
	app.SendMessage("User", "Hello")
}
