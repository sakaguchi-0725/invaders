package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

type Game struct {
	player *Player
}

type GameObject interface {
	Update() error
	Draw(screen *ebiten.Image)
}

func (g *Game) Update() error {
	return g.player.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Run() error {
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Invaders")

	if err := ebiten.RunGame(g); err != nil {
		return err
	}

	return nil
}

func New() (*Game, error) {
	player, err := NewPlayer()
	if err != nil {
		return nil, err
	}

	return &Game{
		player: player,
	}, nil
}
