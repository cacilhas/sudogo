package ui

import (
	raygui "github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

var mainMenu Scene

type mainMenuType struct {
}

func init() {
	mainMenu = &mainMenuType{}
}

func (menu *mainMenuType) Init() Scene {
	raylib.SetExitKey(raylib.KeyEscape)
	return menu
}

func (menu *mainMenuType) Render() Scene {
	width := float32(viper.GetInt("width"))
	height := float32(viper.GetInt("height"))

	titleWidth := height * 0.867
	buttonWidth := width / 3
	bigFontSize := int64(height / 7.5)
	buttonFontSize := int64(height / 10)

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

	raygui.SetStyleProperty(raygui.GlobalTextFontsize, buttonFontSize)
	btX := (width - buttonWidth) / 2
	btY := float32(bigFontSize)*1.5 + height/30
	btHeight := float32(buttonFontSize) * 1.2
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"Very Easy",
	) {
		return mainMenu // TODO: start easy game
	}

	btY += float32(buttonFontSize) * 1.25
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"Easy",
	) {
		return mainMenu // TODO: start easy game
	}

	btY += float32(buttonFontSize) * 1.25
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"Medium",
	) {
		return mainMenu // TODO: start medium game
	}

	btY += float32(buttonFontSize) * 1.25
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"Hard",
	) {
		return mainMenu // TODO: start hard game
	}

	btY += float32(buttonFontSize) * 1.25
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"Fiendish",
	) {
		return mainMenu // TODO: start fiendish game
	}

	return menu
}
