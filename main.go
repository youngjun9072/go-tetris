package main

import (
	termbox "github.com/nsf/termbox-go"
	"go-tetris/block"
	"time"
)

const boardFillColor = termbox.ColorBlack
const boardLineColor = termbox.ColorWhite

const boardWidth = 12
const boardHeight = 20

var board [boardHeight][boardWidth]int

var existBlock = 0

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

func drawBoard(b *block.Block) {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if b.Piece[0][y][x] == -1 {
				board[y+b.Y][x+b.X] = b.Piece[0][y][x]
			}
		}
	}

	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			if board[y][x] == 1 {
				drawBlock(x*4, y*2, boardLineColor)
			} else if board[y][x] == -1 {
				drawBlock(x*4, y*2, termbox.ColorYellow)
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
	block.InitBlock()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	var b *block.Block
loop:
	for {
		if existBlock == 0 {
			b = block.NewBlock()
			existBlock = 1
		}
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyCtrlX:
					break loop
				case termbox.KeyArrowLeft:
					b.MoveToLeft()
				case termbox.KeyArrowRight:
					b.MoveToRight()
				case termbox.KeyArrowDown:
					b.MoveToDown()
				case termbox.KeyArrowUp:
					b.Rotate()
				}
			}
		default:
			initBoard()
			drawBoard(b)
			termbox.Flush()
			time.Sleep(10 * time.Nanosecond)

		}
	}
}
