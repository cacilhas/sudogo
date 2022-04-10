package ui

import (
	"image/color"

	"github.com/cacilhas/sudogo/sudoku"
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

type gameplayType struct {
	sudoku.Game
}

var colours [10]color.RGBA = [10]color.RGBA{
	raylib.Black,
	raylib.Red,
	raylib.Orange,
	raylib.Yellow,
	raylib.Green,
	raylib.SkyBlue,
	raylib.Blue,
	raylib.Violet,
	raylib.Pink,
	raylib.DarkGray,
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
	player.move()
	xOffset := (width - boardSize) / 2
	yOffset := (height - boardSize) / 2
	drawBoard(xOffset, yOffset, boardSize)
	drawGame(xOffset, yOffset, boardSize/9, gameplay.Game)
	player.render(xOffset, yOffset, boardSize/9)

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

func drawGame(sx, sy, cellSize int32, game sudoku.Game) {
	smallSize := cellSize / 3
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			ix := sx + int32(x)*cellSize
			iy := sy + int32(y)*cellSize
			cell := game.Get(x, y)
			if cell.IsSet() {
				raylib.DrawCircle(
					ix+cellSize/2,
					iy+cellSize/2,
					float32(cellSize)/2,
					colours[cell.Value()],
				)
			} else {
				for i := 1; i <= 9; i++ {
					if cell.Candidate(i) {
						ix2 := int32(i-1) % 3
						iy2 := int32(i-1) / 3
						raylib.DrawCircle(
							ix+int32(ix2)*smallSize+smallSize/2,
							iy+int32(iy2)*smallSize+smallSize/2,
							float32(smallSize)/2,
							colours[i],
						)
					}
				}
			}
		}
	}
}
