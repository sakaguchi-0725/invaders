package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 440
)

type Game struct {
	player  *Player
	enemies []*Enemy
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

	for _, enemy := range g.enemies {
		if err := enemy.Update(); err != nil {
			return err
		}
	}

	for _, bullet := range g.bullets {
		if err := bullet.Update(); err != nil {
			return err
		}
	}

	// 画面外の敵を削除
	activeEnemies := []*Enemy{}
	for _, enemy := range g.enemies {
		if !enemy.IsOutOfScreen() {
			activeEnemies = append(activeEnemies, enemy)
		}
	}
	g.enemies = activeEnemies

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
	for _, enemy := range g.enemies {
		enemy.Draw(screen)
	}

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

	enemies, err := makeEnemies()
	if err != nil {
		return nil, err
	}

	g.enemies = enemies

	return g, nil
}

func makeEnemies() ([]*Enemy, error) {
	initialY := -50.0
	spacing := 35.0
	startX := 20.0

	var enemies []*Enemy
	for i := 0; i < 8; i++ {
		x := startX + float64(i)*spacing
		enemy, err := NewEnemy(x, initialY, Normal)
		if err != nil {
			return nil, err
		}
		enemies = append(enemies, enemy)
	}

	return enemies, nil
}
