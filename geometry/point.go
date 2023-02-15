package geometry

import (
	"fabric_sim/helper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Point struct {
	X, Y         float64
	PrevX, PrevY float64
	pinned       bool
	using        bool
}

func (p *Point) Update() {
	mX, mY := ebiten.CursorPosition()

	if helper.Distance(float64(mX), float64(mY), p.X, p.Y) < 5 && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		p.using = true
	}

	if p.using {
		p.X = float64(mX)
		p.Y = float64(mY)
	}

	if helper.Distance(float64(mX), float64(mY), p.X, p.Y) < 5 && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		p.using = false
	}

	if p.pinned {
		return
	}

	forceX := 0.
	forceY := 0.05
	accelerationX := forceX
	accelerationY := forceY

	prevPosX := p.X
	prevPosY := p.Y

	p.X = p.X*2 - p.PrevX + accelerationX*((1/60)*(1/60))
	p.Y = p.Y*2 - p.PrevY + accelerationY*((1/60)*(1/60))

	p.PrevX = prevPosX
	p.PrevY = prevPosY

}

func (p *Point) Draw(screen *ebiten.Image) {
	radius := 3.
	clr := colornames.Orange
	if p.pinned {
		radius = 4.
		clr = colornames.Green
	}
	ebitenutil.DrawCircle(screen, p.X, p.Y, radius, clr)
}

func (p *Point) Pin() {
	p.pinned = true
}
