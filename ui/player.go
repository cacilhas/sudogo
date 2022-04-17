package ui

import (
	"image/color"

	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/spf13/viper"
)

var player struct {
	x, y     int32
	render   func(int32, int32, int32)
	render2D func(int32, int32, int32)
	render3D func(int32, int32, int32)
	move     func()
	goLeft   func()
	goRight  func()
	goUp     func()
	goDown   func()
}

var colour2D = raylib.Brown
var colour3D = color.RGBA{0, 128, 255, 120}

func init() {
	player.x = 4
	player.y = 4

	player.render = func(x, y, z int32) {
		if viper.GetViper().GetBool("3d_rendering") {
			player.render3D(x, y, z)
		} else {
			player.render2D(x, y, z)
		}
	}

	player.render3D = func(x, y, z int32) {
		raylib.DrawSphere(
			raylib.NewVector3(float32(player.x-4+x), float32(4-player.y+y), float32(z)/2),
			0.6,
			colour3D,
		)
	}

	player.render2D = func(x, y, size int32) {
		x0 := x + player.x*size
		y0 := y + player.y*size
		raylib.DrawRectangleLines(x0, y0, size, size, colour2D)
		raylib.DrawRectangleLines(x0+1, y0+1, size-2, size-2, colour2D)
		raylib.DrawRectangleLines(x0+2, y0+2, size-4, size-4, colour2D)
	}

	player.move = func() {
		if !(raylib.IsKeyDown(raylib.KeyLeftControl) || raylib.IsKeyDown(raylib.KeyRightControl)) {
			if raylib.IsKeyPressed(raylib.KeyLeft) || raylib.IsKeyPressed(raylib.KeyA) || raylib.IsKeyPressed(raylib.KeyH) {
				player.goLeft()
			}
			if raylib.IsKeyPressed(raylib.KeyRight) || raylib.IsKeyPressed(raylib.KeyD) || raylib.IsKeyPressed(raylib.KeyL) {
				player.goRight()
			}
			if raylib.IsKeyPressed(raylib.KeyUp) || raylib.IsKeyPressed(raylib.KeyW) || raylib.IsKeyPressed(raylib.KeyK) {
				player.goUp()
			}
			if raylib.IsKeyPressed(raylib.KeyDown) || raylib.IsKeyPressed(raylib.KeyS) || raylib.IsKeyPressed(raylib.KeyJ) {
				player.goDown()
			}
		}
	}

	player.goLeft = func() {
		player.x--
		if player.x < 0 {
			player.x += 9
		}
	}

	player.goRight = func() {
		player.x = (player.x + 1) % 9
	}

	player.goUp = func() {
		player.y--
		if player.y < 0 {
			player.y += 9
		}
	}

	player.goDown = func() {
		player.y = (player.y + 1) % 9
	}
}
