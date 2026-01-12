package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ GameObject = (*Player)(nil)

const (
	playerAcceleration = 0.1
	playerMaxSpeed     = 3.0
	playerFriction     = 0.9
)

type Player struct {
	*Sprite
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(p.Scale, p.Scale)
	opts.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Image, &opts)
}

func (p *Player) Update() error {
	// 左右移動（速度を変更）
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.VelocityX += playerAcceleration
		if p.VelocityX > playerMaxSpeed {
			p.VelocityX = playerMaxSpeed
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.VelocityX -= playerAcceleration
		if p.VelocityX < -playerMaxSpeed {
			p.VelocityX = -playerMaxSpeed
		}
	} else {
		// キーが押されていない時は減速
		p.VelocityX *= playerFriction
		if p.VelocityX > -0.1 && p.VelocityX < 0.1 {
			p.VelocityX = 0
		}
	}

	p.X += p.VelocityX

	// 左端チェック
	if p.X < 0 {
		p.X = 0
		p.VelocityX = 0
	}

	// 右端チェック
	if p.X > ScreenWidth-p.Width {
		p.X = ScreenWidth - p.Width
		p.VelocityX = 0
	}

	return nil
}

func NewPlayer() (*Player, error) {
	image, _, err := ebitenutil.NewImageFromFile("./assets/img/sprite_ship.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load image: %w", err)
	}

	const scale = 0.5

	// 初期位置: 画面下部中央
	shipWidth := float64(image.Bounds().Dx()) * scale
	shipHeight := float64(image.Bounds().Dy()) * scale
	initialX := (ScreenWidth - shipWidth) / 2
	initialY := ScreenHeight - shipHeight - 5

	return &Player{
		&Sprite{
			Image:  image,
			Scale:  scale,
			Width:  shipWidth,
			Height: shipHeight,
			X:      initialX,
			Y:      initialY,
		},
	}, nil
}
