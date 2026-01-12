package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Image     *ebiten.Image
	Scale     float64
	Width     float64
	X, Y      float64
	VelocityX float64
}

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

func (g *Game) Update() error {
	// 加速度の設定
	const acceleration = 0.1
	const maxSpeed = 3.0
	const friction = 0.9

	// 左右移動（速度を変更）
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.VelocityX += acceleration
		if g.VelocityX > maxSpeed {
			g.VelocityX = maxSpeed
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.VelocityX -= acceleration
		if g.VelocityX < -maxSpeed {
			g.VelocityX = -maxSpeed
		}
	} else {
		// キーが押されていない時は減速
		g.VelocityX *= friction
		if g.VelocityX > -0.1 && g.VelocityX < 0.1 {
			g.VelocityX = 0
		}
	}

	g.X += g.VelocityX

	// 左端チェック
	if g.X < 0 {
		g.X = 0
		g.VelocityX = 0
	}

	// 右端チェック
	if g.X > ScreenWidth-g.Width {
		g.X = ScreenWidth - g.Width
		g.VelocityX = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(g.Scale, g.Scale)
	opts.GeoM.Translate(g.X, g.Y)
	screen.DrawImage(g.Image, &opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Invaders")

	image, _, err := ebitenutil.NewImageFromFile("./assets/img/sprite_ship.png")
	if err != nil {
		log.Fatal(err)
	}

	const scale = 0.5

	// 初期位置: 画面下部中央
	shipWidth := float64(image.Bounds().Dx()) * scale
	initialX := (ScreenWidth - shipWidth) / 2
	initialY := ScreenHeight - float64(image.Bounds().Dy())*scale - 5

	err = ebiten.RunGame(&Game{
		Image: image,
		Scale: scale,
		Width: shipWidth,
		X:     initialX,
		Y:     initialY,
	})

	if err != nil {
		log.Fatal(err)
	}
}
