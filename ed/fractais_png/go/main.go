package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func espiralColorido(pen *Pen, dist float64) {
	if dist < 1 {
		return // acaba recursao
	}
	//r, g, b := rand.Float64(), rand.Float64(), rand.Float64()
	pen.SetRGB(ri(0, 255), ri(0, 255), ri(0, 255))
	pen.Walk(dist) // anda quatidade
	pen.Right(90)    // gira em 90 graus
	espiralColorido(pen, dist-7) // recursao
}

func main() {
	pen := NewPen(500, 500)  
	pen.SetPosition(0, 0) // fazer o quadrado
	pen.FillSquare(500, 500)
	pen.SetRGB(0, 0, 0)     // cor
	pen.SetPosition(0, 0) // onde a caenta vai comecar
	espiralColorido(pen, 500)

	pen.SavePNG("espiralColorido.png")
	fmt.Println("PNG file created successfully.")
}
