package handlers

import (
	"gaccess/external/data/repositories"
	"gaccess/internal/utils"
	"gaccess/logic/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRedirect(c *gin.Context, database *gorm.DB) {
	redirectRepository := repositories.SiteRedirectRepository{Database: database}
	accessInfoRepository := repositories.AccesInfoRepository{Database: database}
	redirect := redirectRepository.FindByCode(c.Param("code"))

	location := utils.GetGeolocation(c.ClientIP())
	accessInfoRepository.Create(&entities.AccessInfo{SiteRedirectId: redirect.Id, Ip: c.ClientIP(), City: location.City, State: location.RegionName, Country: location.Country})
	c.Redirect(302, redirect.Redirect)
}
