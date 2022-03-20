package block

import (
	"math/rand"
	"time"
)

const numOfRotate = 4
const pieceWidth = 4
const pieceHeight = 4

type Block struct {
	X     int
	Y     int
	Rot   int
	Piece [numOfRotate][pieceWidth][pieceHeight]int
}

// ㅁ, ㅣ, ㄹ, ㅗ, ㄴ
const numOfBlock = 5

var blocks [numOfBlock]Block

func InitBlock() {
	blocks[0].Piece = [4][4][4]int{
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, -1, -1},
			{0, 0, -1, -1}},
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, -1, -1},
			{0, 0, -1, -1}},
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, -1, -1},
			{0, 0, -1, -1}},
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, -1, -1},
			{0, 0, -1, -1}},
	}
	blocks[1].Piece = [4][4][4]int{
		{{0, 0, -1, 0},
			{0, 0, -1, 0},
			{0, 0, -1, 0},
			{0, 0, -1, 0}},
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{-1, -1, -1, -1},
			{0, 0, 0, 0}},
	}
	blocks[2].Piece = [4][4][4]int{
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, -1, 0, 0},
			{0, -1, -1, -1}},
		{{0, 0, 0, 0},
			{0, 0, 0, -1},
			{0, 0, 0, -1},
			{0, 0, -1, -1}},
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, -1, -1, -1},
			{0, 0, 0, -1}},
		{{0, 0, 0, 0},
			{0, 0, -1, -1},
			{0, 0, -1, 0},
			{0, 0, -1, 0}},
	}
	blocks[3].Piece = [4][4][4]int{
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, -1, -1, 0},
			{0, 0, -1, -1}},
		{{0, 0, 0, 0},
			{0, -1, 0, 0},
			{0, -1, -1, 0},
			{0, 0, -1, 0}},
	}
	blocks[4].Piece = [4][4][4]int{
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, -1, 0},
			{0, -1, -1, -1}},
		{{0, 0, 0, 0},
			{0, 0, 0, -1},
			{0, 0, -1, -1},
			{0, 0, 0, -1}},
		{{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, -1, -1, -1},
			{0, 0, -1, 0}},
		{{0, 0, 0, 0},
			{0, 0, -1, 0},
			{0, 0, -1, -1},
			{0, 0, -1, 0}},
	}

}
func NewBlock() *Block {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	b := blocks[r.Intn(5)]
	b.Rot = 0
	b.X = 0
	b.Y = 0

	return &b
}

func (b *Block) MoveToLeft() {
	if b.X > 1 {
		b.X--
	}
}

func (b *Block) MoveToRight() {
	if b.X < 14 {
		b.X++
	}
}

func (b *Block) MoveToDown() {
	b.Y++
}

func (b *Block) Rotate() {
	if b.Rot == 3 {
		b.Rot = 0
	} else {
		b.Rot++
	}

}
