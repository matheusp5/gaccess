package http

import (
	"gaccess/internal/http/handlers"
	"gaccess/internal/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

const PORT uint16 = 3030
const IP string = "0.0.0.0"

var newLogger logger.Logger = logger.NewLogger()

func getGinDefault() *gin.Engine {
	r := gin.Default()
	return r
}

func runServer(r *gin.Engine) {
	newLogger.Info("Server running on " + IP + ":" + strconv.Itoa(int(PORT)))
	err := r.Run(IP + ":" + strconv.Itoa(int(PORT)))
	if err != nil {
		newLogger.Error("An error occurred while running HTTP server: " + err.Error())
		panic("")
	}
}

func registerRoutes(r *gin.Engine, database *gorm.DB) {
	r.GET("/r/:code", func(c *gin.Context) {
		handlers.GetRedirect(c, database)
	})

	r.POST("/", func(c *gin.Context) {
		handlers.CreateRedirect(c, database)
	})

	r.GET("/c/:id", func(c *gin.Context) {
		handlers.GetInfos(c, database)
	})

	newLogger.Info("Routes registered with success")
}

func InitHttpServer(database *gorm.DB) {
	r := getGinDefault()

	registerRoutes(r, database)
	runServer(r)
}
