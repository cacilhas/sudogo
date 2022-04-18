package ui

import (
	"fmt"
	"image/color"
	"time"

	"github.com/cacilhas/rayframe"
	raygui "github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

var helpMessage string = `HELP [%s]

F1 :: Show this help
F2 :: Toogle 3D rendering
Control + F :: Toggle fullscreen
Arrow keys / WASD / HJKL :: Move
ESC :: Back to main menu / Exit
Numeric keys :: Toggle candidates
Shift + Numeric keys :: Set cell value`

// Grant all required interfaces are implemented
type Help interface {
	rayframe.InitScene
	rayframe.ExitKeyScene
	rayframe.BackgroundScene
	rayframe.UpdateScene
	rayframe.RendererScene2D
}

type helpType struct {
	*rayframe.RayFrame
	previous rayframe.Scene
}

func showHelp(previous rayframe.Scene) Help {
	return &helpType{previous: previous}
}

func (help *helpType) Init(frame *rayframe.RayFrame) {
	help.RayFrame = frame
}

func (help helpType) ExitKey() int32 {
	return 0
}

func (help helpType) OnKeyEscape() rayframe.Scene {
	return help.previous
}

func (help helpType) Background() color.RGBA {
	return raylib.RayWhite
}

func (help helpType) Update(dt time.Duration) rayframe.Scene {
	update(dt)
	return help
}

func (help helpType) Render2D() rayframe.Scene {
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
		fmt.Sprintf(helpMessage, viper.GetString("version")),
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
