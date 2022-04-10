package main

import (
	"math/rand"
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	readSettings()
	defer saveSettings()

	raylib.InitWindow(
		viper.GetInt32("width"),
		viper.GetInt32("height"),
		"Sudogo",
	)

	raylib.SetWindowMinSize(800, 600)
	raylib.SetWindowState(raylib.FlagWindowResizable)
	raylib.SetTargetFPS(24)

	raylib.CloseWindow()
}
