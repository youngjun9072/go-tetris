package utils

import "github.com/nsf/termbox-go"

const CellWidth = 4
const CellHeight = 2

func DrawBlock(x, y int, color termbox.Attribute) {
	for i := x; i < x+CellWidth; i++ {
		for j := y; j < y+CellHeight; j++ {
			termbox.SetCell(i, j, ' ', color, color)
		}
	}
}
