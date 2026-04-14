package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func arvore(pen *Pen, dist float64) {
	if dist < 11{
		return
	}
	angulo := 21.2
	fator := 0.778

	pen.Walk(dist)
	pen.Right(angulo)
	arvore(pen, dist*fator)
	pen.Left(2 * angulo)
	arvore(pen, dist*fator)

	pen.Right(angulo)
	pen.Walk(-dist)
}

func main() {
	pen := NewPen(750, 900)
	pen.SetHeading(90)
	pen.SetPosition(0, 0)
	pen.FillSquare(800, 1000)
	pen.SetPosition(375, 860)
	pen.SetRGB(255, 255, 255)
	arvore(pen, 160)

	pen.SavePNG("arvore.png")
	fmt.Println("PNG file created successfully.")
}
