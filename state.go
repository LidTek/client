package main

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Enter(g *Game)
	Leave(g *Game)
	Update(g *Game) error
	Draw(g *Game, screen *ebiten.Image)
}
