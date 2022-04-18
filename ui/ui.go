package ui

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

func update(dt time.Duration) {
	dealWithPointer(dt)
	fullscreen := raylib.IsWindowFullscreen()
	if getFullscreen() != fullscreen {
		raylib.ToggleFullscreen()
	}
}

func getFullscreen() bool {
	fullscreen := viper.GetBool("fullscreen")
	control := raylib.IsKeyDown(raylib.KeyLeftControl) || raylib.IsKeyDown(raylib.KeyRightControl)
	if raylib.IsKeyPressed(raylib.KeyF) && control {
		fullscreen = !fullscreen
		viper.Set("fullscreen", fullscreen)
	}
	return fullscreen
}
