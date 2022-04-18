package main

import (
	"math/rand"
	"time"

	"github.com/cacilhas/rayframe"
	"github.com/cacilhas/sudogo/ui"
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	readSettings()
	defer saveSettings()

	camera := raylib.NewCamera3D(
		raylib.NewVector3(0, 0, 10),
		raylib.NewVector3(0, 0, 0),
		raylib.NewVector3(0, 1, 0),
		50.0,
		raylib.CameraPerspective,
	)

	frame := &rayframe.RayFrame{
		Camera:   &camera,
		FPS:      24,
		OnRezise: onResize,
	}

	frame.Init(
		viper.GetInt("width"),
		viper.GetInt("height"),
		"Sudogo",
	)
	raylib.SetWindowMinSize(800, 600)
	raylib.SetWindowState(raylib.FlagWindowResizable)

	frame.Mainloop(ui.MainMenu)
}

func onResize(width, height int) {
	if !raylib.IsWindowFullscreen() {
		viper.Set("width", width)
		viper.Set("height", height)
	}
}
