package repo

import(
	"app/MyGoTemplate/db/entities"
	"app/MyGoTemplate/logger"

	"gorm.io/gorm"
)

var Localization *LocalizationRepo

func init() {
	//to use singleton instance
	Localization = new(LocalizationRepo)
}

type LocalizationRepo struct{}

func (u *LocalizationRepo) Create(db *gorm.DB, localization entities.Localization) (error) {
	if err := db.Create(&localization).Error; err != nil {
		logger.ErrorLog("An error occured while creating localization - Create - localizationRepo.go", localization, err.Error())
		return err
	}
	logger.TransactionLog("Created", localization)
	return nil
}

func (u *LocalizationRepo) First(db *gorm.DB, resource string, language string) (entities.Localization, error) {
	var localization entities.Localization
	if err := db.Where("resource = ? and language = ?", resource, language).First(&localization).Error; err != nil {
		logger.ErrorLog("An error occured while getting first localization - First - localizationRepo.go ", "resource = ? and language = ?", resource, language, err.Error())
		return localization, err
	}
	logger.TransactionLog("First", localization)
	return localization, nil
}

func (u *LocalizationRepo) Update(db *gorm.DB, resource string, message string, language string) (error) {
	if err := db.Model(&entities.Localization{}).Where("resource = ? and language = ?", resource, language).Update("message", message).Error; err != nil {
		logger.ErrorLog("An error occured while updating localization - Update - localizationRepo.go ", "resource = ? and language = ?", resource, language, err.Error())
		return err
	}
	logger.TransactionLog("Update", "resource = ? and language = ?", resource, language, "message", message)
	return nil
}

func (u *LocalizationRepo) Delete(db *gorm.DB, resource string, language string) (error) {
	if err := db.Where("resource = ? and language = ?", resource, language).Delete(&entities.Localization{}).Error; err != nil {
		logger.ErrorLog("An error occured while deleting localization - Delete - localizationRepo.go ", "resource = ? and language = ?", resource, language, err.Error())
		return err
	}
	logger.TransactionLog("Delete", "resource = ? and language = ?", resource, language)
	return nil
}

func (u *LocalizationRepo) GetAll(db *gorm.DB) ([]entities.Localization, error) {
	var localizations []entities.Localization
	if err := db.Find(&localizations).Error; err != nil {
		logger.ErrorLog("An error occured while getting all localizations - GetAll - localizationRepo.go ", err.Error())
		return localizations, err
	}
	logger.TransactionLog("GetAll", localizations)
	return localizations, nil
}