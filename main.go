package main

import (
	"flag"
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {
	const (
		emptyCell = ' '
		ballCell  = '⚾'

		speed = time.Second / 20
	)

	var (
		cx, cy        int
		width, height int
		vx, vy        = 1, 1

		cell rune
	)

	flag.IntVar(&width, "width", 20, "Ширина окна")
	flag.IntVar(&height, "height", 10, "Высота окна")
	flag.Parse()

	board := make([][]bool, width)
	for x := range board {
		board[x] = make([]bool, height)
	}

	buf := make([]rune, 0, width*height*2+height)

	screen.Clear()

	for {
		screen.MoveTopLeft()

		cx += vx
		cy += vy

		if cx >= width-1 || cx <= 0 {
			vx *= -1
		}

		if cy >= height-1 || cy <= 0 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[cx][cy] = true
		buf = buf[:0]

		for y := range board[0] {
			for x := range board {
				cell = emptyCell
				if board[x][y] {
					cell = ballCell
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		fmt.Print(string(buf))
		//break
		time.Sleep(speed)

	}
}
