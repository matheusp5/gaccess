package handlers

import (
	"gaccess/external/data/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func GetInfos(c *gin.Context, database *gorm.DB) {
	accessInfoRepository := repositories.AccesInfoRepository{Database: database}
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)
	infos := accessInfoRepository.FindByRedirectId(id)
	c.JSON(200, gin.H{
		"infos": infos,
	})
}
