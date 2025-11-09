package hello

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error { return nil }
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello from Android!")
}
func (g *Game) Layout(w, h int) (int, int) { return 320, 240 }

func NewGame() (*Game, error) {
	game := &Game{}
	return game, nil
}
