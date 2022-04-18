package ui

import (
	"image/color"
	"os"
	"time"

	"github.com/cacilhas/rayframe"
	"github.com/cacilhas/sudogo/sudoku"
	raygui "github.com/gen2brain/raylib-go/raygui"
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

type gameplayType struct {
	sudoku.Game
	*rayframe.RayFrame
	enable3D bool
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

func startGameplay(level sudoku.Level) interface{} {
	player.x = 4
	player.y = 4
	game, _ := sudoku.NewGame(level)
	return &gameplayType{game, nil, false}
}

func loadGameplay(input string) (interface{}, error) {
	player.x = 4
	player.y = 4
	if game, err := sudoku.NewGame(input); err == nil {
		return &gameplayType{game, nil, false}, nil
	} else {
		return nil, err
	}
}

func (gameplay *gameplayType) Init(frame *rayframe.RayFrame) {
	gameplay.RayFrame = frame
	raylib.SetExitKey(0)
}

func (gameplay *gameplayType) Background() color.RGBA {
	return raylib.RayWhite
}

func (gameplay *gameplayType) Update(dt time.Duration) interface{} {
	gameplay.enable3D = viper.GetViper().GetBool("3d_rendering")
	if raylib.IsKeyPressed(raylib.KeyF2) {
		gameplay.enable3D = !gameplay.enable3D
		viper.Set("3d_rendering", gameplay.enable3D)
	}
	if raylib.IsKeyPressed(raylib.KeyEscape) {
		return MainMenu
	}
	if raylib.IsKeyPressed(raylib.KeyF1) {
		return showHelp(gameplay)
	}
	update(dt)
	player.move()
	play(gameplay.Game)
	return gameplay
}

//------------------------------------------------------------------------------
// 2D rendering

func (gameplay *gameplayType) Render2D() interface{} {
	if !gameplay.enable3D {
		xOffset, yOffset, boardSize := gameplay.getOffset()
		drawBoard2D(xOffset, yOffset, boardSize)
		drawGame2D(xOffset, yOffset, boardSize/9, gameplay.Game)

		if !raylib.IsCursorHidden() && raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
			cellClicked2D(xOffset, yOffset, boardSize/9)
		}
		player.render(xOffset, yOffset, boardSize/9)
	}
	gameplay.showGameOver()
	return gameplay
}

func drawBoard2D(x, y, size int32) {
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
	for ly := y; ly <= y+cell*9; ly += cell {
		raylib.DrawLine(x, ly, x+cell*9, ly, raylib.Black)
	}
}

func drawGame2D(sx, sy, cellSize int32, game sudoku.Game) {
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
}

func cellClicked2D(x, y, cellSize int32) {
	cellX := (int32(mouseLastPosition.X) - x) / cellSize
	cellY := (int32(mouseLastPosition.Y) - y) / cellSize
	if cellX >= 0 && cellX < 9 && cellY >= 0 && cellY < 9 {
		player.x = cellX
		player.y = cellY
	}
}

//------------------------------------------------------------------------------
// 3D rendering

func (gameplay *gameplayType) Render3D() interface{} {
	if !gameplay.enable3D {
		return gameplay
	}
	drawBoard3D()
	drawGame3D(gameplay.Game)
	xOffset, yOffset, boardSize := gameplay.getOffset()
	if !raylib.IsCursorHidden() && raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
		cellClicked2D(xOffset, yOffset, boardSize/9)
	}
	player.render(0, 0, 1)
	return gameplay
}

func drawBoard3D() {
	clr1 := raylib.White
	clr2 := raylib.LightGray
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			clr := clr2
			if ((x/3)+(y/3))%2 == 0 {
				clr = clr1
			}
			raylib.DrawCube(
				raylib.NewVector3(float32(x-4), float32(y-4), 0),
				0.98, 0.98, 0.01, clr,
			)
		}
		raylib.DrawCube(
			raylib.NewVector3(0, float32(y)-3.5, 0),
			9, 0.02, 1, raylib.Black,
		)
	}
	raylib.DrawCube(
		raylib.NewVector3(0, -4.5, 0),
		9, 0.02, 1, raylib.Black,
	)
	for x := 0; x < 10; x++ {
		raylib.DrawCube(
			raylib.NewVector3(float32(x)-4.5, 0, 0),
			0.02, 9, 1, raylib.Black,
		)
	}
}

func drawGame3D(game sudoku.Game) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			cell := game.Get(x, y)
			cellCenter := raylib.NewVector3(float32(x-4), float32(4-y), 0.5)
			if cell.IsSet() {
				raylib.DrawSphere(cellCenter, 0.48, colours[cell.Value()])
			} else {
				for i := 1; i <= 9; i++ {
					if cell.Candidate(i) {
						raylib.DrawSphere(
							raylib.NewVector3(
								cellCenter.X+float32((i-1)%3-1)*0.35,
								cellCenter.Y+float32((i-1)/3-1)*0.35,
								0.2,
							),
							0.15, colours[i],
						)
					}
				}
			}
		}
	}
}

//------------------------------------------------------------------------------
// Other functions

func play(game sudoku.Game) {
	control := raylib.IsKeyDown(raylib.KeyLeftControl) || raylib.IsKeyDown(raylib.KeyRightControl)
	shift := raylib.IsKeyDown(raylib.KeyLeftShift) || raylib.IsKeyDown(raylib.KeyRightShift)
	if raylib.IsKeyPressed(raylib.KeyS) && control {
		saveCurrentBoard(game.String())
	}

	if raylib.IsKeyPressed(raylib.KeySpace) {
		game.Autofill()
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

func (gameplay *gameplayType) showGameOver() {
	if gameplay.GameOver() {
		width := gameplay.WindowSize.X
		height := gameplay.WindowSize.Y
		raygui.SetStyleColor(raygui.LabelTextColor, raylib.Maroon)
		raygui.SetStyleProperty(raygui.GlobalTextFontsize, int64(width)/20)
		raygui.Label(
			raylib.Rectangle{X: 0, Y: 0, Width: float32(width), Height: float32(height)},
			"Board Solved!!",
		)
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

func (gameplay *gameplayType) getOffset() (int32, int32, int32) {
	width := int32(gameplay.WindowSize.X)
	height := int32(gameplay.WindowSize.Y)
	boardSize := int32(height/9) * 9
	if width < boardSize {
		boardSize = int32(width/9) * 9
	}
	return (width - boardSize) / 2, (height - boardSize) / 2, boardSize
}
