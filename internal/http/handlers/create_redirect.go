package handlers

import (
	"gaccess/external/data/repositories"
	"gaccess/logic/dto"
	"gaccess/logic/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRedirect(c *gin.Context, database *gorm.DB) {
	redirectRepository := repositories.SiteRedirectRepository{Database: database}

	dto := dto.CreateRedirectDTO{}
	c.BindJSON(&dto)

	redirect := entities.SiteRedirect{Redirect: dto.Redirect}
	redirectRepository.Create(&redirect)

	host := c.Request.Host

	c.JSON(201, gin.H{
		"status": 201,
		"url":    host + "/r/" + redirect.Code,
		"id":     redirect.Id,
	})
}
