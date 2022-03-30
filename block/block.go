package block

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

const numOfRotate = 4
const PieceWidth = 4
const PieceHeight = 4

type Block struct {
	X           int
	Y           int
	Rot         int
	NumOfRotate int
	Color       termbox.Attribute
	Piece       [numOfRotate][PieceWidth][PieceHeight]int
}

// ㅁ, ㅣ, ㄹ, ㅗ, ㄴ
const numOfBlock = 5

var blocks [numOfBlock]Block

func InitBlock() {
	blocks[0].Piece = [4][4][4]int{
		{{1, 1, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
	}
	blocks[0].NumOfRotate = 1
	blocks[0].Color = termbox.ColorBlue

	blocks[1].Piece = [4][4][4]int{
		{{2, 0, 0, 0},
			{2, 0, 0, 0},
			{2, 0, 0, 0},
			{2, 0, 0, 0}},
		{{2, 2, 2, 2},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
	}
	blocks[1].NumOfRotate = 2
	blocks[1].Color = termbox.ColorYellow

	blocks[2].Piece = [4][4][4]int{
		{{3, 0, 0, 0},
			{3, 3, 3, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{{0, 3, 0, 0},
			{0, 3, 0, 0},
			{3, 3, 0, 0},
			{0, 0, 0, 0}},
		{{3, 3, 3, 0},
			{0, 0, 3, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{{3, 3, 0, 0},
			{3, 0, 0, 0},
			{3, 0, 0, 0},
			{0, 0, 0, 0}},
	}
	blocks[2].NumOfRotate = 4
	blocks[2].Color = termbox.ColorCyan

	blocks[3].Piece = [4][4][4]int{
		{{4, 4, 0, 0},
			{0, 4, 4, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{{0, 4, 0, 0},
			{4, 4, 0, 0},
			{4, 0, 0, 0},
			{0, 0, 0, 0}},
	}
	blocks[3].NumOfRotate = 2
	blocks[3].Color = termbox.ColorLightMagenta

	blocks[4].Piece = [4][4][4]int{
		{{0, 5, 0, 0},
			{5, 5, 5, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{{0, 5, 0, 0},
			{5, 5, 0, 0},
			{0, 5, 0, 0},
			{0, 0, 0, 0}},
		{{5, 5, 5, 0},
			{0, 5, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0}},
		{{5, 0, 0, 0},
			{5, 5, 0, 0},
			{5, 0, 0, 0},
			{0, 0, 0, 0}},
	}
	blocks[4].NumOfRotate = 4
	blocks[4].Color = termbox.ColorDarkGray

}
func NewBlock() *Block {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	b := blocks[r.Intn(5)]
	b.Rot = r.Intn(b.NumOfRotate)
	b.X = 10 / 2
	b.Y = 1

	return &b
}

func (b *Block) MoveToLeft() {
	if b.X > 1 {
		b.X--
	}
}

func (b *Block) MoveToRight() {
	if b.X < 12 {
		b.X++
	}
}

func (b *Block) MoveToDown() {
	if b.Y < 19 {
		b.Y++
	}
}

func (b *Block) Rotate(idx int) {
	b.Rot = b.Rot + idx
	if b.Rot >= b.NumOfRotate {
		b.Rot = 0
	}

	if b.Rot == -1 {
		b.Rot = b.NumOfRotate - 1
	}

}
