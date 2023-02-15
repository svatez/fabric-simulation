package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type Stick struct {
	start, end *Point
	length     float64
}

func (s *Stick) Update() {
	diffX := s.start.X - s.end.X
	diffY := s.start.Y - s.end.Y
	diffFactor := (s.length - Length(diffX, diffY)) / Length(diffX, diffY) * 0.5
	offsetX := diffX * diffFactor
	offsetY := diffY * diffFactor

	s.start.X += offsetX
	s.start.Y += offsetY
	s.end.X -= offsetX
	s.end.Y -= offsetY

}

func (s *Stick) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, s.start.X, s.start.Y, s.end.X, s.end.Y, colornames.Orange)
}
