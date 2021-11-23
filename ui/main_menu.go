package ui

import (
	"github.com/cacilhas/sudogo/sudoku"
	raygui "github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

type mainMenu struct {
}

func NewMainMenu() Scene {
	return &mainMenu{}
}

func (m *mainMenu) Init() Scene {
	raylib.SetExitKey(raylib.KeyEscape)
	return m
}

func (m *mainMenu) Render() Scene {
	width := viper.GetInt32("width")
	height := viper.GetInt32("height")
	bigFont := int64(float32(width) / 7.5)
	titleWidth := float32(bigFont) * 4
	if bigFont > 120 {
		bigFont = 120
	}

	y := float32(height) / 30
	menuFont := int64(float32(height) / 12.5)
	raygui.SetStyleColor(raygui.LabelTextColor, raylib.LightGray)
	raygui.SetStyleProperty(raygui.GlobalTextFontsize, bigFont)
	raygui.Label(
		raylib.Rectangle{
			X:      (float32(width) - titleWidth) / 2,
			Y:      y,
			Width:  titleWidth,
			Height: float32(bigFont),
		},
		"Sudogo",
	)

	boxSize := float32(height) / 15
	boxX := (float32(width) - boxSize - float32(menuFont)*4) / 2

	raygui.SetStyleColor(raygui.LabelTextColor, raylib.RayWhite)
	raygui.SetStyleProperty(raygui.GlobalTextFontsize, menuFont)

	y += float32(bigFont) * 1.5

	for _, hs := range []sudoku.Hardship{sudoku.EASY, sudoku.MEDIUM, sudoku.HARD, sudoku.FIENDISH} {
		if raygui.CheckBox(
			raylib.Rectangle{
				X:      boxX,
				Y:      y,
				Width:  boxSize,
				Height: boxSize,
			},
			viper.GetInt("hardship") == int(hs),
		) {
			viper.Set("hardship", int(hs))
		}
		raygui.Label(
			raylib.Rectangle{
				X:      boxX + float32(menuFont)*1.22,
				Y:      y + boxSize/2 - float32(menuFont)/2,
				Width:  float32(menuFont) * 4,
				Height: float32(menuFont),
			},
			hs.String(),
		)
		y += float32(menuFont) * 1.25
	}

	y += float32(menuFont) / 4
	if raygui.Button(
		raylib.Rectangle{
			X:      boxX + boxSize,
			Y:      y,
			Width:  float32(menuFont) * 4,
			Height: float32(menuFont),
		},
		"Play",
	) {
		// TODO: start game
	}

	return m
}
