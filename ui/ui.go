package ui

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

var windowWidth int
var windowHeight int

func Mainloop() {
	windowWidth = raylib.GetScreenWidth()
	windowHeight = raylib.GetScreenHeight()
	scene := mainMenu.Init()

	for !raylib.WindowShouldClose() {
		fullscreen := raylib.IsWindowFullscreen()
		shouldBeFullscreen := getFullscreen()

		if shouldBeFullscreen && !fullscreen {
			raylib.ToggleFullscreen()
		} else if !shouldBeFullscreen && fullscreen {
			raylib.ToggleFullscreen()
		}
		fullscreen = shouldBeFullscreen

		if raylib.IsWindowResized() {
			windowWidth = raylib.GetScreenWidth()
			windowHeight = raylib.GetScreenHeight()
			if !fullscreen {
				viper.Set("width", windowWidth)
				viper.Set("height", windowHeight)
			}
		}

		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.RayWhite)

		scene = scene.Render()

		raylib.EndDrawing()
		time.Sleep(time.Millisecond * 42)
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
