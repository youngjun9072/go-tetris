package block

const numOfRotate = 4
const pieceWidth = 4
const pieceHeight = 4

type Block struct {
	x     int
	y     int
	rot   int
	piece [numOfRotate][pieceWidth][pieceHeight]int
}

// ㅁ, ㅣ, ㄹ, ㅗ, ㄴ
const numOfBlock = 5

var blocks [numOfBlock]Block

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
func NewBlock() *Block {
	b := new(Block)
	b.rot = 0
	b.x = 0
	b.y = 0

	return b
}

func (b *Block) MoveToLeft() {

}

func (b *Block) MoveToRight() {

}

func (b *Block) MoveToDown() {

}

func (b *Block) Rotate() {
}
