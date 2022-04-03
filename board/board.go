package board

import (
	"go-tetris/block"
)

const Width = 13
const Height = 20

const boundary = -1
const field = 0

var Board [Height][Width]int

func Init() {
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if x == 0 || y == 0 || x == Width-1 || y == Height-1 {
				Board[y][x] = boundary
			} else {
				Board[y][x] = field
			}
		}
	}
}

func PutOn(b *block.Block) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				Board[b.Y+y][b.X+x] = b.Piece[b.Rot][y][x]
			}
		}
	}
}

func Erase(b *block.Block) {
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if b.Piece[b.Rot][y][x] != 0 {
				Board[b.Y+y][b.X+x] = 0
			}
		}
	}
}

func removeLine(start int) {
	for y := start; y > 1; y-- {
		for x := 1; x < Width; x++ {
			Board[y][x] = Board[y-1][x]
		}
	}
}

func RemoveLines() {
	y := Height - 2
	for {
		check := true

		if y == 1 {
			break
		}
		for x := 1; x < Width; x++ {
			if Board[y][x] == 0 {
				check = false
			}
		}

		if check {
			removeLine(y)
		} else {
			y--
		}
	}
}
