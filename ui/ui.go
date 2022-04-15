package ui

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

var windowWidth int
var windowHeight int
var mouseStop time.Duration
var hideMouse time.Duration = 1_500 * time.Millisecond
var mouseLastPosition raylib.Vector2

func Mainloop() {
	windowWidth = raylib.GetScreenWidth()
	windowHeight = raylib.GetScreenHeight()
	scene := mainMenu.Init()
	lastTick := time.Now()
	mouseStop = time.Duration(0)
	mouseLastPosition = raylib.GetMousePosition()

	for !raylib.WindowShouldClose() {
		tick := time.Now()
		dealWithPointer(tick.Sub(lastTick))
		lastTick = tick
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

func dealWithPointer(dt time.Duration) {
	if raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
		raylib.ShowCursor()
		mouseStop = time.Duration(0)
		return
	}
	position := raylib.GetMousePosition()
	if position.X != mouseLastPosition.X || position.Y != mouseLastPosition.Y {
		raylib.ShowCursor()
		mouseLastPosition = position
		mouseStop = time.Duration(0)
		return
	}
	mouseStop += dt
	if mouseStop >= hideMouse {
		raylib.HideCursor()
	}
}
