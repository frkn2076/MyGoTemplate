package repositories

import(
	"app/MyGoTemplate/db"
	"app/MyGoTemplate/db/entities"
	s "app/MyGoTemplate/session"
)

type Localization struct{}

func (u *Localization) Create(resource string, message string, language string) (err error) {
	if err := db.GormDB.Where("resource = ? and language = ?", resource, language).First(&entities.Login{}).Error; err != nil {
		localization := entities.Localization{Resource: resource, Message: message, Language: language}
		db.GormDB.Create(&localization)
		return nil
	} else {
		return err
	}
}

func (u *Localization) First(resource string) (message string, err error) {
	language := s.SessionGet("language")
	var localization Localization
	result := db.GormDB.Where("resource = ? and language = ?", resource, language).First(&localization)
	if result.Error != nil {
		return "", result.Error
	} else {
		return localization, nil
	}
}