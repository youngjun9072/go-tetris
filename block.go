package main

const numOfRotate = 4
const pieceWidth = 4
const pieceHeight = 4

type block struct {
	x     int
	y     int
	rot   int
	piece [numOfRotate][pieceWidth][pieceHeight]int
}

// ㅁ, ㅣ, ㄹ, ㅗ, ㄴ
const numOfBlock = 5

var blocks [numOfBlock]block

func InitBlock() {
	blocks[0].piece = [4][4][4]int{
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

}
func NewBlock() *block {
	b := new(block)
	b.rot = 0
	b.x = 0
	b.y = 0
	for y := b.y; y < 4; y++ {
		for x := b.x; x < 4; x++ {
			if blocks[0].piece[b.rot][y][x] == -1 {
				board[y][x] = blocks[0].piece[b.rot][y][x]
			}

		}
	}
	existBlock = 1
	return b
}

func (b *block) moveToLeft() {

}

func (b *block) moveToRight() {

}

func (b *block) moveToDown() {

}

func (b *block) rotate() {
}
