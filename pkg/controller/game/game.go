package game

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var _ Handler = (*HandlerImpl)(nil)

type Handler interface {
	SetNewGame(gameID string, xSize, ySize int) (*PlayData, error)
	PlaceStone(gameID string, x, y int) error
	Observe(gameID string) (result [][]int, err error)
	GetPlayData(gameID string) (*PlayData, error)
	GetResultBoard(gameID string) ([][]int, error)
}

type HandlerImpl struct {
	PlayDataMap map[string]*PlayData
}

func NewHandlerImpl() *HandlerImpl {
	playDataMap := make(map[string]*PlayData)
	return &HandlerImpl{
		PlayDataMap: playDataMap,
	}
}

func (h *HandlerImpl) SetNewGame(gameID string, xSize, ySize int) (*PlayData, error) {
	logrus.Debug("controller.game.SetNewGame Called")
	h.PlayDataMap[gameID] = NewPlayData(xSize, ySize)
	logrus.Debug("controller.game.SetNewGame Finished")
	return h.PlayDataMap[gameID], nil
}

func (h *HandlerImpl) PlaceStone(gameID string, x, y int) error {
	logrus.Debug("controller.game.PlaceStone Called")
	playData, err := h.GetPlayData(gameID)
	if err != nil {
		logrus.Error("game_id", gameID, "not found, ", err)
		return err
	}

	if err := playData.SetStone(x, y); err != nil {
		logrus.Errorf("[%d, %d] is not empty in gameID %s", x, y, gameID)
		return err
	}
	logrus.Debug("controller.game.PlaceStone Finished")
	return nil
}

func (h *HandlerImpl) Observe(gameID string) (result [][]int, err error) {
	logrus.Debug("controller.game.Observe Called")
	playData, err := h.GetPlayData(gameID)
	if err != nil {
		return nil, err
	}
	result = playData.Observe()
	logrus.Debug("controller.game.Observe Finished")
	return result, nil
}

func (h *HandlerImpl) GetPlayData(gameID string) (*PlayData, error) {
	logrus.Debug("controller.game.GetPlayData Called")
	data, ok := h.PlayDataMap[gameID]
	if !ok {
		return nil, fmt.Errorf("404 not found")
	}
	logrus.Debug("controller.game.GetPlayData Finished")
	return data, nil
}

func (h *HandlerImpl) GetResultBoard(gameID string) ([][]int, error) {
	logrus.Debug("controller.game.GetResultBoard Called")
	data, ok := h.PlayDataMap[gameID]
	if !ok {
		return nil, fmt.Errorf("404 not found")
	}
	logrus.Debug("controller.game.GetResultBoard Finished")
	return data.ResultBoard.GetCurResult(), nil
}
