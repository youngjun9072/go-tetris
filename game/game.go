package game

import (
	"github.com/nsf/termbox-go"
	"go-tetris/block"
	"go-tetris/board"
	"go-tetris/utils"
)

var StartGame bool
var Score uint
var Level uint
var KeyLock bool

func InitGame() {
	Score = 0
	Level = 0
	KeyLock = true
	board.Init()
	block.Init()
}

func DrawBoard() {
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			utils.DrawBlock(x*4, y*2, utils.GetColor(board.Board[y][x]))
		}
	}
}

func CheckRightCollid(b *block.Block) bool {
	for x := 3; x >= 0; x-- {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				if board.Board[b.Y+y][b.X+x+1] != 0 {
					return true
				}
			}
		}
	}
	return false
}

func CheckLeftCollid(b *block.Block) bool {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				if board.Board[b.Y+y][b.X+x-1] != 0 {
					return true
				}
			}
		}
	}
	return false
}

func CheckDownCollid(b *block.Block) bool {
	for x := 3; x >= 0; x-- {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				if board.Board[b.Y+y+1][b.X+x] != 0 {
					return true
				}
			}
		}
	}
	return false
}

func isRotatable(b *block.Block) bool {
	b.Rotate(1)
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if b.X+x > 11 || board.Board[b.Y+y][b.X+x] != 0 && b.Piece[b.Rot][y][x] != 0 {
				b.Rotate(-1)
				return false
			}
		}
	}
	return true
}

func IsEnd() bool {
	if board.Board[1][10/2] != 0 {
		return true
	}
	return false
}

func KeyHandler(ev termbox.Event, b *block.Block) bool {
	if ev.Type == termbox.EventKey {
		switch ev.Key {
		case termbox.KeyCtrlX:
			return false

		case termbox.KeyCtrlS:
			if !StartGame {
				StartGame = true
				InitGame()
			}

		case termbox.KeyArrowLeft:
			if StartGame || !KeyLock {
				board.Erase(b)
				if !CheckLeftCollid(b) {
					b.MoveToLeft()
				}
				board.PutOn(b)
			}
			return true
		case termbox.KeyArrowRight:
			if StartGame || !KeyLock {
				board.Erase(b)
				if !CheckRightCollid(b) {
					b.MoveToRight()
				}
				board.PutOn(b)
			}
			return true
		case termbox.KeyArrowDown:
			if StartGame || !KeyLock {
				board.Erase(b)
				collid := CheckDownCollid(b)
				if !collid {
					b.MoveToDown()
				}
				board.PutOn(b)
				if collid {
					board.RemoveLines()
					b = nil
				}
			}
			return true
		case termbox.KeyArrowUp:
			if StartGame || !KeyLock {
				board.Erase(b)
				if !isRotatable(b) {

				}
				board.PutOn(b)
			}
			return true
		case termbox.KeySpace:
			if StartGame || !KeyLock {
				KeyLock = true
				for {
					board.Erase(b)
					if !CheckDownCollid(b) {
						b.MoveToDown()
					} else {
						break
					}
					board.PutOn(b)
				}
				b = nil
				board.RemoveLines()
			}
			return true
		}
	}
	return true
}
