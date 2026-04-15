package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func circulos(pen *Pen, dist float64) {
	if dist < 6 {
		return
	}
	angulo := 60.0
	fator := 0.339

	pen.DrawCircle(dist)
	for range 6 {
		pen.SetLineWidth(0.99)
		pen.Right(angulo)

		pen.Up()
		pen.Walk(dist)
		pen.Down()

		pen.DrawCircle(dist * fator)
		circulos(pen, dist*fator)

		pen.Up()
		pen.Walk(-dist)
		pen.Down()
	}
}

func main() {
	pen := NewPen(800, 800)
	pen.SetHeading(90)
	pen.SetPosition(0, 0)
	pen.FillSquare(800, 800)

	pen.SetPosition(400, 420)
	pen.SetRGB(250, 250, 250)
	circulos(pen, 280)

	pen.SavePNG("circulos.png")
	fmt.Println("PNG file created successfully.")
}
