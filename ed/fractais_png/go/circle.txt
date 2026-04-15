package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func circulo(pen *Pen, raio float64, nivel int) {
	if nivel == 0 {
		return
	}

	pen.DrawCircle(raio)
	
	for i := 0; i < 6; i++ {
		pen.Right(60)

		pen.Up()
		pen.Walk(raio)
		pen.Down()

		pen.Right(90)

		circulo(pen, raio*0.35, nivel-1)

		pen.Left(90)
		pen.Up()
		pen.Walk(-raio)
		pen.Down()
	}
}

func main() {
	pen := NewPen(1200, 1200)
	pen.SetPosition(600, 600)
	circulo(pen, 320, 5)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
