package utils

import "github.com/nsf/termbox-go"

const CellWidth = 4
const CellHeight = 2

var ColorMap = map[int]termbox.Attribute{
	-1: termbox.ColorWhite,
	0:  termbox.ColorBlack,
	1:  termbox.ColorBlue,
	2:  termbox.ColorYellow,
	3:  termbox.ColorCyan,
	4:  termbox.ColorLightMagenta,
	5:  termbox.ColorDarkGray,
}

func DrawBlock(x, y int, color termbox.Attribute) {
	for i := x; i < x+CellWidth; i++ {
		for j := y; j < y+CellHeight; j++ {
			termbox.SetCell(i, j, ' ', color, color)
		}
	}
}

func GetColor(i int) termbox.Attribute {
	return ColorMap[i]
}
