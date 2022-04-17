package ui

import (
	raygui "github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
)

var helpMessage string = `HELP

F1 :: Show this help
F2 :: Toogle 3D rendering
Control + F :: Toggle fullscreen
Arrow keys / WASD / HJKL :: Move
ESC :: Back to main menu / Exit
Numeric keys :: Toggle candidates
Shift + Numeric keys :: Set cell value`

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

	width := float32(windowWidth)
	height := float32(windowHeight)

	titleWidth := height * 0.867
	textWidth := width * 0.75
	bigFontSize := int64(height / 7.5)
	textFontSize := bigFontSize / 3

	raygui.SetStyleColor(raygui.LabelTextColor, raylib.Black)
	raygui.SetStyleProperty(raygui.GlobalTextFontsize, textFontSize)
	raygui.Label(
		raylib.Rectangle{
			X:      (width - textWidth) / 2,
			Y:      -height / 5, // FIXME: why not zero?
			Width:  textWidth,
			Height: height,
		},
		helpMessage,
	)

	raygui.SetStyleColor(raygui.LabelTextColor, raylib.DarkBlue)
	raygui.SetStyleProperty(raygui.GlobalTextFontsize, bigFontSize)
	raygui.Label(
		raylib.Rectangle{
			X:      (width - titleWidth) / 2,
			Y:      height / 30,
			Width:  titleWidth,
			Height: float32(bigFontSize),
		},
		"Sudogo",
	)

	return help
}
