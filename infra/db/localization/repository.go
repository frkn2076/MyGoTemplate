package localization

import (
	"app/MyGoTemplate/infra/logger"

	"gorm.io/gorm"
)

func NewRepository() *Repository {
	return new(Repository)
}

type Repository struct{}

func (u *Repository) Create(db *gorm.DB, localization Entity) error {
	if err := db.Create(&localization).Error; err != nil {
		logger.ErrorLog("An error occured while creating localization - Create - localizationRepo.go", localization, err.Error())
		return err
	}
	logger.TransactionLog("Created", localization)
	return nil
}

func (u *Repository) First(db *gorm.DB, resource string, language string) (Entity, error) {
	var localization Entity
	if err := db.Where("resource = ? and language = ?", resource, language).First(&localization).Error; err != nil {
		logger.ErrorLog("An error occured while getting first localization - First - localizationRepo.go ",
		 "resource = ? and language = ?", resource, language, err.Error())
		return localization, err
	}
	logger.TransactionLog("First", localization)
	return localization, nil
}

func (u *Repository) Update(db *gorm.DB, resource string, message string, language string) error {
	if err := db.Model(&Entity{}).Where("resource = ? and language = ?", resource, language).Update("message", message).Error; err != nil {
		logger.ErrorLog("An error occured while updating localization - Update - localizationRepo.go ",
		 "resource = ? and language = ?", resource, language, err.Error())
		return err
	}
	logger.TransactionLog("Update", "resource = ? and language = ?", resource, language, "message", message)
	return nil
}

func (u *Repository) Delete(db *gorm.DB, resource string, language string) error {
	if err := db.Where("resource = ? and language = ?", resource, language).Delete(&Entity{}).Error; err != nil {
		logger.ErrorLog("An error occured while deleting localization - Delete - localizationRepo.go ",
		 "resource = ? and language = ?", resource, language, err.Error())
		return err
	}
	logger.TransactionLog("Delete", "resource = ? and language = ?", resource, language)
	return nil
}

func (u *Repository) GetAll(db *gorm.DB) ([]Entity, error) {
	var localizations []Entity
	if err := db.Find(&localizations).Error; err != nil {
		logger.ErrorLog("An error occured while getting all localizations - GetAll - localizationRepo.go ", err.Error())
		return localizations, err
	}
	logger.TransactionLog("GetAll", localizations)
	return localizations, nil
}
