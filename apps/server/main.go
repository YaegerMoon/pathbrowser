package main

import "github.com/YaegerMoon/pathbrowser/server/app"

func main() {
	a := app.New("localhost", "8080")
	a.Start()
}
