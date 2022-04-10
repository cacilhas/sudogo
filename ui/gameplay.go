package ui

import (
	"github.com/cacilhas/sudogo/sudoku"
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

type gameplayType struct {
	sudoku.Game
}

func startGameplay(level sudoku.Level) Scene {
	return &gameplayType{sudoku.NewGame(level)}
}

func (gameplay *gameplayType) Init() Scene {
	raylib.SetExitKey(0)
	return gameplay
}

func (gameplay *gameplayType) Render() Scene {
	if raylib.IsKeyPressed(raylib.KeyEscape) {
		return mainMenu.Init()
	}

	width := viper.GetInt32("width")
	height := viper.GetInt32("height")
	boardSize := height
	if width < boardSize {
		boardSize = width
	}
	//game := gameplay.Game
	xOffset := (width - boardSize) / 2
	yOffset := (height - boardSize) / 2
	drawBoard(xOffset, yOffset, boardSize)

	return gameplay
}

func drawBoard(x, y, size int32) {
	blockSize := size / 3
	cell := size / 9
	clr1 := raylib.RayWhite
	clr2 := raylib.LightGray
	raylib.DrawRectangle(x, y, blockSize, blockSize, clr1)
	raylib.DrawRectangle(x+blockSize, y, blockSize, blockSize, clr2)
	raylib.DrawRectangle(x+2*blockSize, y, blockSize, blockSize, clr1)
	raylib.DrawRectangle(x, y+blockSize, blockSize, blockSize, clr2)
	raylib.DrawRectangle(x+blockSize, y+blockSize, blockSize, blockSize, clr1)
	raylib.DrawRectangle(x+2*blockSize, y+blockSize, blockSize, blockSize, clr2)
	raylib.DrawRectangle(x, y+2*blockSize, blockSize, blockSize, clr1)
	raylib.DrawRectangle(x+blockSize, y+2*blockSize, blockSize, blockSize, clr2)
	raylib.DrawRectangle(x+2*blockSize, y+2*blockSize, blockSize, blockSize, clr1)

	for lx := x; lx <= x+cell*9; lx += cell {
		raylib.DrawLine(lx, y, lx, y+cell*10, raylib.Black)
	}
	for ly := y; ly <= y+cell*10; ly += cell {
		raylib.DrawLine(x, ly, x+cell*9, ly, raylib.Black)
	}
}
