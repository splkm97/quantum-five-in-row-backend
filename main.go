package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
	"quantum-five-in-row-backend/pkg/route"
	"time"
)

var (
	servePort = "8080"
	upgrader  = websocket.Upgrader{}
)

func main() {
	// TODO: log level 설정 만들기
	logrus.SetLevel(logrus.DebugLevel)

	e := gin.Default()
	log := logrus.New()
	e.Use(ginlogrus.Logger(log), gin.Recovery())
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	route.SetRouteRules(e)
	e.Run(":" + servePort)
}
