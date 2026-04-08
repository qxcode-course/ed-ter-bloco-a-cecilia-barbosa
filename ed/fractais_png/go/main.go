package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func embua(pen *Pen, dist float64) {
	if dist < 1 {
		return // acaba recursao
	}
	//r, g, b := rand.Float64(), rand.Float64(), rand.Float64()
	//pen.SetRGB(r, g, b)
	pen.SetLineWidth(dist / 30)
	pen.Walk(dist) // anda quatidade
	dist *= 0.97
	pen.Right(90)    // gira em 90 graus
	embua(pen, dist) // recursao
}

func main() {
	pen := NewPen(500, 500)   // criar o tamanho da imagem
	pen.SetRGB(134, 0, 0)     // cor
	pen.SetPosition(100, 100) // onde a caenta vai comecar
	embua(pen, 300)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
