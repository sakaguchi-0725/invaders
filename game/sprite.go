package game

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Image                *ebiten.Image
	Scale                float64
	Width, Height        float64
	X, Y                 float64
	VelocityX, VelocityY float64
}
