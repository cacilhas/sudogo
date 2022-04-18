package ui

import (
	"image/color"
	"time"

	"github.com/cacilhas/rayframe"
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
	*rayframe.RayFrame
	previous interface{}
}

func showHelp(previous interface{}) interface{} {
	return &helpType{previous: previous}
}

func (help *helpType) Init(frame *rayframe.RayFrame) {
	help.RayFrame = frame
	raylib.SetExitKey(0)
}

func (help *helpType) Background() color.RGBA {
	return raylib.RayWhite
}

func (help *helpType) Update(dt time.Duration) interface{} {
	if raylib.IsKeyPressed(raylib.KeyEscape) {
		return help.previous
	}
	update(dt)
	return help
}

func (help *helpType) Render2D() interface{} {
	width := float32(help.WindowSize.X)
	height := float32(help.WindowSize.Y)

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
