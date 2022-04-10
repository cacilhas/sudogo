package ui

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

func Mainloop() {
	scene := mainMenu.Init()

	for !raylib.WindowShouldClose() {
		fullscreen := raylib.IsWindowFullscreen()
		shouldBeFullscreen := viper.GetBool("fullscreen")

		if shouldBeFullscreen && !fullscreen {
			raylib.ToggleFullscreen()
		} else if !shouldBeFullscreen && fullscreen {
			raylib.ToggleFullscreen()
		}
		fullscreen = shouldBeFullscreen
		viper.Set("fullscreen", fullscreen)

		if raylib.IsWindowResized() && !fullscreen {
			viper.Set("width", raylib.GetScreenWidth())
			viper.Set("height", raylib.GetScreenHeight())
		}

		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.RayWhite)

		scene = scene.Render()

		raylib.EndDrawing()
		time.Sleep(time.Millisecond * 42)
	}
}
