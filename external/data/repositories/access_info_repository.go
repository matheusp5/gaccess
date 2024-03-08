package repositories

import (
	"gaccess/internal/utils"
	"gaccess/logic/entities"
	"gorm.io/gorm"
)

type AccesInfoRepository struct {
	Database *gorm.DB
}

func (repository *AccesInfoRepository) GetAll() []entities.AccessInfo {
	var result []entities.AccessInfo
	repository.Database.Find(&result)
	return result
}

func (repository *AccesInfoRepository) Create(accessInfo *entities.AccessInfo) {
	accessInfo.Id = utils.GenerateId()
	repository.Database.Create(&accessInfo)
}

func (repository *AccesInfoRepository) FindById(id uint32) entities.AccessInfo {
	var result entities.AccessInfo
	repository.Database.First(&result, id)
	return result
}

func (repository *AccesInfoRepository) FindByRedirectId(id int) entities.AccessInfo {
	var result entities.AccessInfo
	repository.Database.Find(&result, "site_redirect_id = ?", id)
	return result
}
