package repo

import(
	"app/MyGoTemplate/db/entities"

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
		return err
	}
	return nil
}

func (u *LocalizationRepo) First(db *gorm.DB, resource string, language string) (entities.Localization, error) {
	var localization entities.Localization
	if err := db.Where("resource = ? and language = ?", resource, language).First(&localization).Error; err != nil {
		return localization, err
	}
	return localization, nil
}

func (u *LocalizationRepo) Update(db *gorm.DB, resource string, message string, language string) (error) {
	if err := db.Model(&entities.Localization{}).Where("resource = ? and language = ?", resource, language).Update("message", message).Error; err != nil {
		return err
	}
	return nil
}

func (u *LocalizationRepo) Delete(db *gorm.DB, resource string, language string) (error) {
	if err := db.Where("resource = ? and language = ?", resource, language).Delete(&entities.Localization{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *LocalizationRepo) GetAll(db *gorm.DB) ([]entities.Localization, error) {
	var localizations []entities.Localization
	if err := db.Find(&localizations).Error; err != nil {
		return localizations, err
	}
	return localizations, nil
}