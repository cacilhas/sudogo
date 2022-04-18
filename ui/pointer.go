package ui

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var mouseStop time.Duration = 0
var hideMouse time.Duration = 1_500 * time.Millisecond
var mouseLastPosition raylib.Vector2

func dealWithPointer(dt time.Duration) {
	if raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
		raylib.ShowCursor()
		mouseStop = 0
		return
	}
	position := raylib.GetMousePosition()
	if position.X != mouseLastPosition.X || position.Y != mouseLastPosition.Y {
		raylib.ShowCursor()
		mouseLastPosition = position
		mouseStop = 0
		return
	}
	mouseStop += dt
	if mouseStop >= hideMouse {
		raylib.HideCursor()
	}
}
