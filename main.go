package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math"
)

func Distance(x1, y1, x2, y2 float64) float64 {
	x, y := x2-x1, y2-y1
	return math.Sqrt(x*x + y*y)
}

func Length(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	points [][]*Point
	sticks []*Stick
}

func NewGame() *Game {
	gm := &Game{}

	space := 25

	gm.points = make([][]*Point, 15)
	for i := range gm.points {
		gm.points[i] = make([]*Point, 15)
	}
	lastI := len(gm.points) - 1
	lastJ := len(gm.points[0]) - 1

	for i := 0; i < lastI+1; i++ {
		for j := 0; j < lastJ+1; j++ {
			point := &Point{float64((i + 1) * space), float64(j * space), 10, float64((i + 1) * space), float64(j*space) - 0.05, false, false}

			if j == 0 || j == lastJ || i == 0 || i == lastI {
				point.Pin()
			}
			gm.points[i][j] = point
		}
	}

	for i := 0; i < len(gm.points)-1; i++ {
		for j := 0; j < len(gm.points)-1; j++ {
			stick1 := NewStick(gm.points[i][j], gm.points[i+1][j])
			stick2 := NewStick(gm.points[i][j], gm.points[i][j+1])
			//stick3 := NewStick(gm.points[i][j], gm.points[i+1][j+1])
			//stick4 := NewStick(gm.points[i][j+1], gm.points[i+1][j])
			//gm.sticks = append(gm.sticks, stick3, stick4)
			gm.sticks = append(gm.sticks, stick1, stick2)

		}
	}

	for i := 0; i < len(gm.points)-1; i++ {
		stick1 := NewStick(gm.points[lastI][i], gm.points[lastI][i+1])
		stick2 := NewStick(gm.points[i][lastJ], gm.points[i+1][lastJ])
		gm.sticks = append(gm.sticks, stick1, stick2)
	}

	return gm
}

func (g *Game) Update() error {
	for _, point := range g.points {
		for _, p := range point {
			p.Update()
		}
	}
	for _, stick := range g.sticks {
		stick.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, point := range g.points {
		for _, p := range point {
			p.Draw(screen)
		}
	}
	for _, stick := range g.sticks {
		stick.Draw(screen)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}
