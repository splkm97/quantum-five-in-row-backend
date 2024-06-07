package game

import (
	"fmt"
	"math/rand"
)

type PlayData struct {
	Board        *Board
	ResultBoard  *ResultBoard
	xSize, ySize int
	nextStone    int
}

func NewPlayData(xSize, ySize int) *PlayData {
	return &PlayData{
		Board:       NewBoard(xSize, ySize),
		ResultBoard: NewResultBoard(xSize, ySize),
		xSize:       xSize,
		ySize:       ySize,
		nextStone:   -2,
	}
}

func (pd *PlayData) GetSize() (xSize, ySize int) {
	return pd.xSize, pd.ySize
}

func (pd *PlayData) GetNextStone() int {
	switch pd.nextStone {
	case -2:
		return 2
	case -1:
		return 1
	case 1:
		return -2
	case 2:
		return -1
	}
	return 0
}

func (pd *PlayData) setNextStone() {
	pd.nextStone = pd.GetNextStone()
}

func (pd *PlayData) SetStone(x, y int) error {
	var stone float64
	switch pd.GetNextStone() {
	case -2:
		stone = 0.1
	case -1:
		stone = 0.3
	case 1:
		stone = 0.7
	case 2:
		stone = 0.9
	}
	curBoard := pd.Board.GetCurState()
	xSize, ySize := pd.GetSize()
	if y >= ySize || x >= xSize {
		return fmt.Errorf("input value over array")
	}
	if curBoard[y][x] != 0 {
		return fmt.Errorf("stone already exists")
	}
	pd.Board.SetStone(x, y, stone)
	pd.setNextStone()
	return nil
}

func (pd *PlayData) Observe() [][]int {
	state := pd.Board.GetCurState()
	xSize, ySize := pd.GetSize()
	result := make([][]int, ySize)

	for i := 0; i < ySize; i++ {
		result[i] = make([]int, xSize)
		for j := 0; j < xSize; j++ {
			if state[i][j] == 0 {
				// 빈 칸
				continue
			}
			threshold := rand.Intn(100)
			cur := int(state[i][j] * 100)
			if cur < threshold {
				// 백돌로 결정
				result[i][j] = 1
			} else {
				// 흑돌로 결정
				result[i][j] = -1
			}
		}
	}

	pd.ResultBoard.SetResult(result)

	return result
}
