package geometry

import (
	"fabric_sim/helper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type Stick struct {
	start, end *Point
	length     float64
}

func NewStick(p1 *Point, p2 *Point) *Stick {
	return &Stick{start: p1, end: p2, length: helper.Distance(p1.X, p1.Y, p2.X, p2.Y)}
}

func (s *Stick) Update() {
	diffX := s.start.X - s.end.X
	diffY := s.start.Y - s.end.Y
	diffFactor := (s.length - helper.Length(diffX, diffY)) / helper.Length(diffX, diffY) * 0.5
	offsetX := diffX * diffFactor
	offsetY := diffY * diffFactor

	s.start.X += offsetX
	s.start.Y += offsetY
	s.end.X -= offsetX
	s.end.Y -= offsetY
}

func (s *Stick) Draw(screen *ebiten.Image) {
	curLength := helper.Distance(s.start.X, s.start.Y, s.end.X, s.end.Y)
	clr := color.RGBA{R: 255, G: uint8((curLength / s.length) * 165), A: 255}
	ebitenutil.DrawLine(screen, s.start.X, s.start.Y, s.end.X, s.end.Y, clr)
}
