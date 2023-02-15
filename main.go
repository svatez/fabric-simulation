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

	space := 35

	gm.points = make([][]*Point, 10)
	for i := range gm.points {
		gm.points[i] = make([]*Point, 10)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			point := &Point{float64((i + 1) * space), float64(j * space), 10, float64((i + 1) * space), float64(j*space) - 0.05, false, false}
			if j == 0 || j == 9 || i == 0 || i == 9 {
				point.Pin()
			}
			gm.points[i][j] = point
		}
	}

	for i := 0; i < len(gm.points)-1; i++ {
		for j := 0; j < len(gm.points)-1; j++ {
			stick1 := &Stick{gm.points[i][j], gm.points[i+1][j], float64(space)}
			stick2 := &Stick{gm.points[i][j], gm.points[i][j+1], float64(space)}
			//stick3 := &Stick{gm.points[i][j], gm.points[i+1][j+1], float64(space) * math.Sqrt(2)}
			//stick4 := &Stick{gm.points[i][j+1], gm.points[i+1][j], float64(space) * math.Sqrt(2)}
			//gm.sticks = append(gm.sticks, stick3, stick4)
			gm.sticks = append(gm.sticks, stick1, stick2)

		}
	}

	for i := 0; i < len(gm.points)-1; i++ {
		stick1 := &Stick{gm.points[9][i], gm.points[9][i+1], float64(space)}
		stick2 := &Stick{gm.points[i][9], gm.points[i+1][9], float64(space)}
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
