package repositories

import (
	"gaccess/internal/utils"
	"gaccess/logic/entities"
	"gorm.io/gorm"
)

type SiteRedirectRepository struct {
	Database *gorm.DB
}

func (repository *SiteRedirectRepository) GetAll() []entities.SiteRedirect {
	var siteRedirects []entities.SiteRedirect
	repository.Database.Find(&siteRedirects)
	return siteRedirects
}

func (repository *SiteRedirectRepository) Create(redirect *entities.SiteRedirect) {
	redirect.Id = utils.GenerateId()
	redirect.Code = utils.GenerateCode()
	redirect.Accesses = 0
	repository.Database.Create(&redirect)
}

func (repository *SiteRedirectRepository) FindById(id uint32) entities.SiteRedirect {
	var result entities.SiteRedirect
	repository.Database.First(&result, id)
	return result
}

func (repository *SiteRedirectRepository) FindByCode(code string) entities.SiteRedirect {
	var result entities.SiteRedirect
	repository.Database.First(&result, "code = ?", code)
	return result
}
