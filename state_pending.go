package main

import (
	"image/color"

	"github.com/LidTek/client/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StatePending struct {
	ticks int
}

func (s *StatePending) Enter(g *Game) {
}

func (s *StatePending) Leave(g *Game) {
}

func (s *StatePending) Update(g *Game) error {
	s.ticks++
	return nil
}

func (s *StatePending) Draw(g *Game, screen *ebiten.Image) {
	screen.Fill(color.White)
	opts := &text.DrawOptions{}
	opts.LayoutOptions.SecondaryAlign = text.AlignCenter
	opts.ColorScale.ScaleWithColor(color.Black)

	w, h := text.Measure("Pending  ", assets.FontGroups[assets.DisplayFont].GetFace(text.WeightBold), 0)

	opts.GeoM.Translate(float64(screen.Bounds().Dx())/2-w/2, float64(screen.Bounds().Dy())/2-h/2)

	face := assets.FontGroups[assets.DisplayFont].GetFace(text.WeightBold)

	if s.ticks%80 < 20 {
		text.Draw(screen, "Pending", face, opts)
	} else if s.ticks%80 < 40 {
		text.Draw(screen, "Pending.", face, opts)
	} else if s.ticks%80 < 60 {
		text.Draw(screen, "Pending..", face, opts)
	} else {
		text.Draw(screen, "Pending...", face, opts)
	}
}
