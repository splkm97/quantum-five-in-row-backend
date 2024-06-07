package game

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"quantum-five-in-row-backend/pkg/controller/game"
	"strconv"
)

var handler game.Handler

func init() {
	handler = game.NewHandlerImpl()
}

func SetNewGame(c *gin.Context) {
	logrus.Debug("api.game.SetNewGame api called")
	var body SetNewGameRequest
	err := c.Bind(&body)
	body.GameID = c.Param("game_id")
	logrus.Debug("xSize: ", body.XSize, ", ySize: ", body.YSize)

	if err != nil {
		resp := SetNewGameResponse{
			Status:  "error",
			Message: "cannot bind request body",
		}
		logrus.Error("api.game.SetNewGame api failed, " + resp.Message)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	playData, err := handler.SetNewGame(body.GameID, body.XSize, body.YSize)
	if err != nil {
		resp := SetNewGameResponse{
			Status:  "error",
			Message: fmt.Sprintf("cannot set new game, %s", err.Error()),
		}
		logrus.Error("api.game.SetNewGame api failed, " + resp.Message)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	logrus.Debug("api.game.SetNewGame api success")
	c.JSON(http.StatusOK, SetNewGameResponse{
		Status:   "success",
		Message:  "set new game successfully",
		PlayData: playData.Board.GetCurState(),
	})
}

func PlaceStone(c *gin.Context) {
	logrus.Debug("api.game.PlaceStone api called")
	gameID := c.Param("game_id")
	xStr := c.Query("x")
	yStr := c.Query("y")
	x, err1 := strconv.Atoi(xStr)
	y, err2 := strconv.Atoi(yStr)
	if err1 != nil || err2 != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := handler.PlaceStone(gameID, x, y); err != nil {
		c.Status(http.StatusNotFound)
		// TODO: status not found / status bad request 분기
		return
	}
	playData, err := handler.GetPlayData(gameID)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, playData.Board.GetCurState())
	logrus.Debug("api.game.PlaceStone api success")
}

func Observe(c *gin.Context) {
	logrus.Debug("api.game.Observe api called")
	gameID := c.Param("game_id")
	result, err := handler.Observe(gameID)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, result)
	logrus.Debug("api.game.Observe api success")
}

func GetResult(c *gin.Context) {
	logrus.Debug("api.game.GetResult api called")
	gameID := c.Param("game_id")
	result, err := handler.GetResultBoard(gameID)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, result)
	logrus.Debug("api.game.GetResult api success")
}

func GetPlayData(c *gin.Context) {
	logrus.Debug("api.game.GetPlayData api called")
	gameID := c.Param("game_id")
	playData, err := handler.GetPlayData(gameID)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	logrus.Debug("api.game.GetPlayData api success")
	c.JSON(http.StatusOK, playData.Board.GetCurState())
	return
}
