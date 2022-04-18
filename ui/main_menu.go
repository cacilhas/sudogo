package ui

import (
	"image/color"
	"os"
	"time"

	"github.com/cacilhas/rayframe"
	"github.com/cacilhas/sudogo/sudoku"
	raygui "github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
)

// Grant all required interfaces are implemented
type MainMenuInterface interface {
	rayframe.InitScene
	rayframe.BackgroundScene
	rayframe.UpdateScene
	rayframe.RendererScene2D
}

type MainMenuType struct {
	*rayframe.RayFrame
}

var MainMenu MainMenuInterface = &MainMenuType{}

func (menu *MainMenuType) Init(frame *rayframe.RayFrame) {
	menu.RayFrame = frame
}

func (menu MainMenuType) Background() color.RGBA {
	return raylib.RayWhite
}

func (menu MainMenuType) Update(dt time.Duration) rayframe.Scene {
	if raylib.IsKeyPressed(raylib.KeyF1) {
		return showHelp(menu)
	}
	update(dt)
	return menu
}

func (menu MainMenuType) Render2D() rayframe.Scene {
	width := float32(menu.WindowSize.X)
	height := float32(menu.WindowSize.Y)

	buttonWidth := width * 0.6
	bigFontSize := int64(height / 7.5)
	buttonFontSize := int64(height / 10)

	raygui.SetStyleColor(raygui.LabelTextColor, raylib.DarkBlue)
	raygui.SetStyleColor(raygui.ButtonDefaultTextColor, raylib.Black)
	raygui.SetStyleProperty(raygui.GlobalTextFontsize, bigFontSize)
	raygui.Label(
		raylib.Rectangle{
			X:      0,
			Y:      height / 30,
			Width:  width,
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
		"0. Very Easy",
	) || raylib.IsKeyPressed(raylib.KeyZero) {
		return startGameplay(sudoku.EXTREMELY_EASY)
	}

	btY += float32(buttonFontSize) * 1.25
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"1. Easy",
	) || raylib.IsKeyPressed(raylib.KeyOne) {
		return startGameplay(sudoku.EASY)
	}

	btY += float32(buttonFontSize) * 1.25
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"2. Medium",
	) || raylib.IsKeyPressed(raylib.KeyTwo) {
		return startGameplay(sudoku.MEDIUM)
	}

	btY += float32(buttonFontSize) * 1.25
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"3. Hard",
	) || raylib.IsKeyPressed(raylib.KeyThree) {
		return startGameplay(sudoku.HARD)
	}

	btY += float32(buttonFontSize) * 1.25
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"4. Fiendish",
	) || raylib.IsKeyPressed(raylib.KeyFour) {
		return startGameplay(sudoku.FIENDISH)
	}

	btY += float32(buttonFontSize) * 1.25
	if raygui.Button(
		raylib.Rectangle{
			X:      btX,
			Y:      btY,
			Width:  buttonWidth,
			Height: btHeight,
		},
		"L. Load from File",
	) || raylib.IsKeyPressed(raylib.KeyL) {
		return loadFromFile(menu)
	}

	return menu
}

func loadFromFile(scene rayframe.Scene) rayframe.Scene {
	var fp *os.File
	if aux, err := openFile(); err == nil {
		fp = aux
	} else {
		showError(err)
		return scene
	}
	defer fp.Close()
	var data [4096]byte
	if _, err := fp.Read(data[:]); err != nil {
		showError(err)
		return scene
	}
	if nextScene, err := loadGameplay(string(data[:])); err == nil {
		return nextScene
	} else {
		showError(err)
	}
	return scene
}
