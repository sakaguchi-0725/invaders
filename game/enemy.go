package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ GameObject = (*Enemy)(nil)

type EnemyType string

const (
	Normal EnemyType = "normal"
	Boss   EnemyType = "boss"
)

type Enemy struct {
	*Sprite
	Type EnemyType
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(e.Scale, e.Scale)
	opts.GeoM.Translate(e.X, e.Y)
	screen.DrawImage(e.Image, &opts)
}

func (e *Enemy) Update() error {
	e.Y += e.VelocityY
	return nil
}

func (e *Enemy) IsOutOfScreen() bool {
	return e.Y > ScreenHeight
}

func NewEnemy(x, y float64, enemyType EnemyType) (*Enemy, error) {
	image, _, err := ebitenutil.NewImageFromFile("./assets/img/sprite_1.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load enemy image: %w", err)
	}

	const scale = 0.5

	enemyWidth := float64(image.Bounds().Dx()) * scale
	enemyHeight := float64(image.Bounds().Dy()) * scale

	return &Enemy{
		Sprite: &Sprite{
			Image:     image,
			Scale:     scale,
			Width:     enemyWidth,
			Height:    enemyHeight,
			X:         x,
			Y:         y,
			VelocityY: 0.5,
		},
		Type: enemyType,
	}, nil
}
