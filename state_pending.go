package main

import (
	"image/color"

	"github.com/LidTek/client/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StatePending struct {
	ticks   int
	message TickMessage
}

func (s *StatePending) Enter(g *Game) {
	s.message = TickMessage{
		ticks:    90,
		messages: []string{"Pending", "Pending.", "Pending..", "Pending..."},
	}
}

func (s *StatePending) Leave(g *Game) {
}

func (s *StatePending) Update(g *Game) error {
	s.ticks++
	s.message.Update()
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

	text.Draw(screen, s.message.Message(), face, opts)
}
