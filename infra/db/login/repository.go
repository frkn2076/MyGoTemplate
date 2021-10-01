package login

import (
	"app/MyGoTemplate/infra/logger"

	"gorm.io/gorm"
)

func NewRepository() *Repository {
	return new(Repository)
}

type Repository struct{}

func (u *Repository) Create(db *gorm.DB, login Entity) (err error) {
	if err := db.Create(&login).Error; err != nil {
		logger.ErrorLog("An error occured while creating login - Create - loginRepo.go ", login, err.Error())
		return err
	}
	logger.TransactionLog("Create", login)
	return nil
}

func (u *Repository) First(db *gorm.DB, userName string) (Entity, error) {
	var login Entity
	if err := db.Where("user_name = ? ", userName).First(&login).Error; err != nil {
		logger.ErrorLog("An error occured while getting first login - First - loginRepo.go ", "user_name = ? ", userName, err.Error())
		return login, err
	}
	logger.TransactionLog("Create", login)
	return login, nil
}
