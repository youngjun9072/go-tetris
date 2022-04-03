package main

import (
	termbox "github.com/nsf/termbox-go"
	"go-tetris/block"
	"go-tetris/utils"
	"time"
)

const boardFillColor = termbox.ColorBlack
const boardLineColor = termbox.ColorWhite

const boardWidth = 13
const boardHeight = 20

var board [boardHeight][boardWidth]int

var colorMap map[int]termbox.Attribute

func initBoard() {
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			if x == 0 || y == 0 || x == boardWidth-1 || y == boardHeight-1 {
				board[y][x] = -1
			} else {
				board[y][x] = 0
			}
		}
	}

	colorMap = make(map[int]termbox.Attribute)
	colorMap[-1] = termbox.ColorWhite
	colorMap[0] = termbox.ColorBlack
	colorMap[1] = termbox.ColorBlue
	colorMap[2] = termbox.ColorYellow
	colorMap[3] = termbox.ColorCyan
	colorMap[4] = termbox.ColorLightMagenta
	colorMap[5] = termbox.ColorDarkGray
}

func drawBoard() {
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			utils.DrawBlock(x*4, y*2, colorMap[board[y][x]])
		}
	}
}

func checkRightCollid(b *block.Block) bool {
	for x := 3; x >= 0; x-- {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				if board[b.Y+y][b.X+x+1] != 0 {
					return true
				}
			}
		}
	}
	return false
}

func checkLeftCollid(b *block.Block) bool {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				if board[b.Y+y][b.X+x-1] != 0 {
					return true
				}
			}
		}
	}
	return false
}

func checkDownCollid(b *block.Block) bool {
	for x := 3; x >= 0; x-- {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				if board[b.Y+y+1][b.X+x] != 0 {
					return true
				}
			}
		}
	}
	return false
}

func putOn(b *block.Block) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				board[b.Y+y][b.X+x] = b.Piece[b.Rot][y][x]
			}
		}
	}
}

func erase(b *block.Block) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				board[b.Y+y][b.X+x] = 0
			}
		}
	}
}

func isRotatable(b *block.Block) bool {
	b.Rotate(1)
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if b.X+x > 11 || board[b.Y+y][b.X+x] != 0 && b.Piece[b.Rot][y][x] != 0 {
				b.Rotate(-1)
				return false
			}
		}
	}
	return true
}

func removeBlock(start int) {
	for y := start; y > 1; y-- {
		for x := 1; x < boardWidth; x++ {
			board[y][x] = board[y-1][x]
		}
	}
}

func removeBlocks() {
	y := boardHeight - 2
	for {
		check := true

		if y == 1 {
			break
		}
		for x := 1; x < boardWidth; x++ {
			if board[y][x] == 0 {
				check = false
			}
		}

		if check {
			removeBlock(y)
		} else {
			y--
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

	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			if b != nil {
				erase(b)
				collid := checkDownCollid(b)
				if !collid {
					b.MoveToDown()
				}
				putOn(b)
				if collid {
					b = nil
					removeBlocks()
				}
			}
		}
	}()

	ignore := false
loop:
	for {
		if b == nil {
			b = block.NewBlock()
			putOn(b)
		}
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyCtrlX:
					break loop
				case termbox.KeyArrowLeft:
					if ignore {
						break
					}
					erase(b)
					if !checkLeftCollid(b) {
						b.MoveToLeft()
					}
					putOn(b)
				case termbox.KeyArrowRight:
					if ignore {
						break
					}
					erase(b)
					if !checkRightCollid(b) {
						b.MoveToRight()
					}
					putOn(b)
				case termbox.KeyArrowDown:
					if ignore {
						break
					}
					erase(b)
					collid := checkDownCollid(b)
					if !collid {
						b.MoveToDown()
					}
					putOn(b)
					if collid {
						removeBlocks()
						b = nil
					}
				case termbox.KeyArrowUp:
					erase(b)
					if !isRotatable(b) {

					}
					putOn(b)
				case termbox.KeySpace:
					ignore = true
					for {
						erase(b)
						if !checkDownCollid(b) {
							b.MoveToDown()
						} else {
							break
						}
						putOn(b)
					}
					ignore = false
				}
			}
		default:
			drawBoard()
			termbox.Flush()

			time.Sleep(50 * time.Millisecond)

		}
	}
}
