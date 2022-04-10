package ui

import (
	raygui "github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

var helpMessage string = `HELP

F1 :: Show this help
Arrow keys / WASD / HJKL :: move
ESC :: Back to main menu / Exit
Numeric keys :: toggle candidates
Numeric keys + Shift :: set cell value`

type helpType struct {
	previous Scene
}

func showHelp(previous Scene) Scene {
	return &helpType{previous: previous}
}

func (help *helpType) Init() Scene {
	raylib.SetExitKey(0)
	return help
}

func (help *helpType) Render() Scene {
	if raylib.IsKeyPressed(raylib.KeyEscape) {
		return help.previous.Init()
	}

	width := float32(viper.GetInt("width"))
	height := float32(viper.GetInt("height"))

	titleWidth := height * 0.867
	textWidth := width * 0.8
	bigFontSize := int64(height / 7.5)
	textFontSize := bigFontSize / 3

	y := height / 30
	raygui.SetStyleProperty(raygui.GlobalTextFontsize, bigFontSize)
	raygui.Label(
		raylib.Rectangle{
			X:      (width - titleWidth) / 2,
			Y:      y,
			Width:  titleWidth,
			Height: float32(bigFontSize),
		},
		"Sudogo",
	)

	y += float32(bigFontSize) * 1.5
	raygui.SetStyleProperty(raygui.GlobalTextFontsize, textFontSize)
	raygui.Label(
		raylib.Rectangle{
			X:      (width - textWidth) / 2,
			Y:      0,
			Width:  textWidth,
			Height: height - y,
		},
		helpMessage,
	)

	return help
}
