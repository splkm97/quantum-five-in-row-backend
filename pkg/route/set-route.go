package route

import (
	gameapi "quantum-five-in-row-backend/pkg/api/game"
	statapi "quantum-five-in-row-backend/pkg/api/stat"

	"github.com/gin-gonic/gin"
)

func SetRouteRules(e *gin.Engine) {
	setGamesAPIRule(e)
	setStatsAPIRule(e)
}

func setGamesAPIRule(e *gin.Engine) {
	gameAPI := e.Group("/api/game")

	gameAPI.POST("/:game_id", gameapi.SetNewGame)
	gameAPI.GET("/:game_id/board", gameapi.GetPlayData)

	gameAPI.POST("/:game_id/stone", gameapi.PlaceStone)

	gameAPI.GET("/:game_id/result", gameapi.GetResult)
	gameAPI.POST("/:game_id/result", gameapi.Observe)
}

func setStatsAPIRule(e *gin.Engine) {
	statAPI := e.Group("/api/stat")

	statAPI.GET("/user/:user_id", statapi.GetUserStats)
}
