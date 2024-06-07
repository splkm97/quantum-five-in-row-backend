package game

// NewBoard 함수는 [xSize by ySize] 크기의 빈 보드판을 생성하는 생성자이다.
func NewBoard(xSize, ySize int) *Board {
	state := make([][]float64, ySize)
	for i := 0; i < ySize; i++ {
		state[i] = make([]float64, xSize)
	}

	for i := 0; i < ySize; i++ {
		for j := 0; j < xSize; j++ {
			state[i][j] = 0
		}
	}
	return &Board{
		curState: state,
		xSize:    xSize,
		ySize:    ySize,
	}
}

// Board 는 양자오목 보드판의 정보를 담는다.
type Board struct {
	// CurState 는 게임의 현재 보드판 상태를 의미함.
	curState     [][]float64
	xSize, ySize int
}

func (b *Board) GetCurState() [][]float64 {
	return b.curState
}

func (b *Board) SetStone(x, y int, stone float64) {
	b.curState[y][x] = stone
}

func (b *Board) RemoveStone(x, y int) {
	b.curState[y][x] = 0
}

type ResultBoard struct {
	curResult    [][]int
	xSize, ySize int
}

func NewResultBoard(xSize int, ySize int) *ResultBoard {
	return &ResultBoard{
		xSize: xSize,
		ySize: ySize,
	}
}

func (rb *ResultBoard) GetCurResult() [][]int {
	return rb.curResult
}

func (rb *ResultBoard) SetResult(target [][]int) {
	rb.curResult = target
}
