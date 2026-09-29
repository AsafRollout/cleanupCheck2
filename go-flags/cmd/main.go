package main

import (
	"go-flags/app"
	"go-flags/fm"
)

// This file shows a simple Go programme that uses the standard feature flags 'EnableNewFeature' and 'TextColour'.

func main() {
	fm.Init()
	app.Run()
}
