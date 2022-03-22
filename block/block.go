package block

import (
	"math/rand"
	"time"
)

const numOfRotate = 4
const pieceWidth = 4
const pieceHeight = 4

type Block struct {
	X           int
	Y           int
	Rot         int
	NumOfRotate int
	Color       termbox.Attribute
	Piece       [numOfRotate][pieceWidth][pieceHeight]int
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
		{{4, 0, 0, 0},
			{4, 4, 0, 0},
			{0, 4, 0, 0},
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
			{5, 0, 0, 0},
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
	b.X = 8 / 2
	b.Y = 1

	return &b
}

func (b *Block) MoveToLeft() {
	if b.X > 1 {
		b.X--
	}
}

func (b *Block) MoveToRight() {
	if b.X < 14-4 {
		b.X++
	}
}

func (b *Block) MoveToDown() {
	b.Y++
}

func (b *Block) Rotate() {
	b.Rot++
	if b.Rot == b.NumOfRotate {
		b.Rot = 0
	}

}
