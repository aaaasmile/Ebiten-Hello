package main

import (
	"myapp/hello"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := hello.NewGame()
	if err != nil {
		panic(err)
	}
	ebiten.RunGame(game)
}
