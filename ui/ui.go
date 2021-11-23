package ui

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

func Mainloop() {
	darkCyan := raylib.NewColor(0x00, 0x52, 0x52, 0xff)
	scene := NewMainMenu().Init()

	for !raylib.WindowShouldClose() {
		if raylib.IsWindowResized() {
			viper.Set("width", raylib.GetScreenWidth())
			viper.Set("height", raylib.GetScreenHeight())
		}

		raylib.BeginDrawing()
		raylib.ClearBackground(darkCyan)
		scene = scene.Render()
		raylib.EndDrawing()

		time.Sleep(time.Millisecond * 42)
	}
}
