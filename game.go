package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	bigFont   *text.GoTextFace
	state     State
	nextState State
}

func NewGame() *Game {
	g := &Game{}
	g.NextState(&StatePending{})
	return g
}

func (g *Game) NextState(state State) {
	g.nextState = state
}

func (g *Game) Update() error {
	if g.state != nil {
		return g.state.Update(g)
	}
	if g.nextState != nil {
		if g.state != nil {
			g.state.Leave(g)
		}
		g.state = g.nextState
		g.nextState = nil
		g.state.Enter(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.state != nil {
		g.state.Draw(g, screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
