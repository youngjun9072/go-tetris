package main

import (
	termbox "github.com/nsf/termbox-go"
)

const boardFillColor = termbox.ColorBlack
const boardLineColor = termbox.ColorWhite

const boardWidth = 12
const boardHeight = 20

var board [boardHeight][boardWidth]int

func initBoard() {
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			if x == 0 || y == 0 || x == boardWidth-1 || y == boardHeight-1 {
				board[y][x] = 1
			} else {
				board[y][x] = 0
			}
		}
	}
}

func drawBoard() {
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			if board[y][x] == 1 {
				drawBlock(x*4, y*2, boardLineColor)
			} else {
				drawBlock(x*4, y*2, boardFillColor)
			}
		}
	}
}

func drawBlock(x, y int, color termbox.Attribute) {
	for i := x; i < x+4; i++ {
		for j := y; j < y+2; j++ {
			termbox.SetCell(i, j, ' ', color, color)
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Flush()

	initBoard()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	var b *block
loop:
	for {
		if existBlock == 0 {
			b = NewBlock()
			existBlock = 1
		}
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyCtrlX:
					break loop
				case termbox.KeyArrowLeft:
					b.moveToLeft()
				case termbox.KeyArrowRight:
					b.moveToRight()
				case termbox.KeyArrowDown:
					b.moveToDown()
				case termbox.KeyArrowUp:
					b.rotate()
				}
			}
		default:
			drawBoard()
			termbox.Flush()
			time.Sleep(1 * time.Second)

		}
	}
}
