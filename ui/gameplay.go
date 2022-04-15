package ui

import (
	"image/color"
	"os"

	"github.com/cacilhas/sudogo/sudoku"
	raygui "github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
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
	raylib.Gray,
}

func startGameplay(level sudoku.Level) Scene {
	player.x = 4
	player.y = 4
	game, _ := sudoku.NewGame(level)
	return &gameplayType{game}
}

func loadGameplay(input string) (Scene, error) {
	player.x = 4
	player.y = 4
	if game, err := sudoku.NewGame(input); err == nil {
		return &gameplayType{game}, nil
	} else {
		return nil, err
	}
}

func (gameplay *gameplayType) Init() Scene {
	raylib.SetExitKey(0)
	return gameplay
}

func (gameplay *gameplayType) Render() Scene {
	if raylib.IsKeyPressed(raylib.KeyEscape) {
		return mainMenu.Init()
	}
	if raylib.IsKeyPressed(raylib.KeyF1) {
		return showHelp(gameplay).Init()
	}

	width := int32(windowWidth)
	height := int32(windowHeight)
	boardSize := height
	if width < boardSize {
		boardSize = width
	}
	player.move()
	play(gameplay.Game)
	xOffset := (width - boardSize) / 2
	yOffset := (height - boardSize) / 2
	drawBoard(xOffset, yOffset, boardSize)
	drawGame(xOffset, yOffset, boardSize/9, gameplay.Game)

	if !raylib.IsCursorHidden() && raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
		cellClicked(xOffset, yOffset, boardSize/9)
	}
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
		raylib.DrawLine(lx, y, lx, y+cell*9, raylib.Black)
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
						iy2 := 2 - int32(i-1)/3
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

	if game.GameOver() {
		width := int64(windowWidth)
		height := int64(windowHeight)
		raygui.SetStyleColor(raygui.LabelTextColor, raylib.Maroon)
		raygui.SetStyleProperty(raygui.GlobalTextFontsize, width/20)
		raygui.Label(
			raylib.Rectangle{X: 0, Y: 0, Width: float32(width), Height: float32(height)},
			"Board Solved!!",
		)
	}
}

func cellClicked(x, y, cellSize int32) {
	cellX := (int32(mouseLastPosition.X) - x) / cellSize
	cellY := (int32(mouseLastPosition.Y) - y) / cellSize
	if cellX >= 0 && cellX < 9 && cellY >= 0 && cellY < 9 {
		player.x = cellX
		player.y = cellY
	}
}

func play(game sudoku.Game) {
	control := raylib.IsKeyDown(raylib.KeyLeftControl) || raylib.IsKeyDown(raylib.KeyRightControl)
	shift := raylib.IsKeyDown(raylib.KeyLeftShift) || raylib.IsKeyDown(raylib.KeyRightShift)
	if raylib.IsKeyPressed(raylib.KeyS) && control {
		saveCurrentBoard(game.String())
	}

	x := int(player.x)
	y := int(player.y)
	for i := int32(0); i <= 9; i++ {
		if raylib.IsKeyPressed(raylib.KeyKp0+i) || raylib.IsKeyPressed(raylib.KeyZero+i) {
			if shift {
				game.Set(x, y, int(i))
			} else {
				game.Toggle(x, y, int(i))
			}
		}
	}
	if raylib.IsKeyPressed(raylib.KeyU) {
		if shift {
			game.Redo()
		} else {
			game.Undo()
		}
	}
}

func saveCurrentBoard(data string) {
	var fp *os.File
	if aux, err := saveFile(); err == nil {
		fp = aux
	} else {
		showError(err)
		return
	}
	defer func() {
		filename := fp.Name()
		fp.Close()
		if _, err := os.Stat(filename); err == nil {
			os.Chmod(filename, 0644)
			showInfo("Board saved to %s", filename)
		} else {
			showError(err)
		}
	}()
	fp.WriteString(data)
}
