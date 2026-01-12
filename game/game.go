package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

type Game struct {
	player  *Player
	bullets []*Bullet
}

type GameObject interface {
	Update() error
	Draw(screen *ebiten.Image)
}

func (g *Game) Update() error {
	if err := g.player.Update(); err != nil {
		return err
	}

	for _, bullet := range g.bullets {
		if err := bullet.Update(); err != nil {
			return err
		}
	}

	// 画面外の弾を削除
	activeBullets := []*Bullet{}
	for _, bullet := range g.bullets {
		if !bullet.IsOutOfScreen() {
			activeBullets = append(activeBullets, bullet)
		}
	}
	g.bullets = activeBullets

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}
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
	g := &Game{
		bullets: []*Bullet{},
	}

	// プレイヤーの射撃コールバックを設定
	onShoot := func(x, y float64) error {
		bullet, err := NewBullet(x, y, PlayerBullet)
		if err != nil {
			return err
		}
		g.bullets = append(g.bullets, bullet)
		return nil
	}

	player, err := NewPlayer(onShoot)
	if err != nil {
		return nil, err
	}

	g.player = player

	return g, nil
}
