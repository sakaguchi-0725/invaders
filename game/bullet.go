package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ GameObject = (*Bullet)(nil)

type BulletType string

const (
	PlayerBullet BulletType = "player"
	EnemyBullet  BulletType = "enemy"
)

type Bullet struct {
	*Sprite
	Type BulletType
}

func (b *Bullet) Update() error {
	b.Y += b.VelocityY
	return nil
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(b.Scale, b.Scale)
	opts.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(b.Image, &opts)
}

func (b *Bullet) IsOutOfScreen() bool {
	return b.Y < -b.Height || b.Y > ScreenHeight
}

func NewBullet(x, y float64, bulletType BulletType) (*Bullet, error) {
	image, _, err := ebitenutil.NewImageFromFile("./assets/img/sprite_0.png")
	if err != nil {
		return nil, fmt.Errorf("failet to load bullet image: %w", err)
	}

	const scale = 0.5
	bulletWidth := float64(image.Bounds().Dx()) * scale
	bulletHeight := float64(image.Bounds().Dy()) * scale

	var velocityY float64
	switch bulletType {
	case PlayerBullet:
		velocityY = -5.0 // 上に飛ぶ
	case EnemyBullet:
		velocityY = 3.0 // 下に飛ぶ
	}

	return &Bullet{
		Sprite: &Sprite{
			Image:     image,
			Scale:     scale,
			Width:     bulletWidth,
			Height:    bulletHeight,
			X:         x - bulletWidth,
			Y:         y,
			VelocityY: velocityY,
		},
		Type: bulletType,
	}, nil
}
