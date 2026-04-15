package main

import (
	"fmt"
)

func desenhaQuadrado(pen *Pen, lado float64) {
	for i := 0; i < 4; i++ {
		pen.Walk(lado)
		pen.Right(90)
	}
}

// func fractal(pen *Pen, lado float64, nivel int) {
// 	if nivel == 0 {
// 		return
// 	}

// 	desenhaQuadrado(pen, lado)
	
// 	for i := 0; i < 4; i++ {
// 		pen.Up()
// 		pen.Walk(lado * 0.75)
// 		pen.Down()
		
// 		fractal(pen, lado*0.5, nivel-1)

// 		pen.Up()
// 		pen.Walk(-lado * 0.75)
// 		pen.Down()

// 		pen.Right(90)
// 	}
// }

func fractal(pen *Pen, lado float64, nivel int) {
	if nivel == 0 {
		return
	}

	desenhaQuadrado(pen, lado)

	passo := lado / 3

	// vai pro canto inferior esquerdo
	pen.Up()
	pen.Walk(-lado / 2)
	pen.Right(90)
	pen.Walk(lado / 2)
	pen.Left(90)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			if !(i == 1 && j == 1) {
				pen.Down()
				fractal(pen, passo, nivel-1)
				pen.Up()
			}

			pen.Walk(passo)
		}

		pen.Walk(-3 * passo)

		pen.Left(90)
		pen.Walk(passo)
		pen.Right(90)
	}

	// 🔥 VOLTA PRO PONTO ORIGINAL (ESSENCIAL)
	pen.Left(90)
	pen.Walk(-lado / 2)
	pen.Right(90)
	pen.Walk(lado / 2)
}

func main() {
	pen := NewPen(1600, 1600)
	pen.SetPosition(400, 400)
	fractal(pen, 800, 2)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
