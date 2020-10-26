package repo

import(
	"app/MyGoTemplate/db/entities"

	"gorm.io/gorm"
)

var Login *LoginRepo

func init() {
	//to use singleton instance
	Login = new(LoginRepo)
}

type LoginRepo struct{}

func (u *LoginRepo) Create(db *gorm.DB, login entities.Login) (err error) {
	if err := db.Create(&login).Error; err != nil {
		logger.ErrorLog("An error occured while creating login - Create - loginRepo.go ", login, err.Error())
		return err
	}
	logger.TransactionLog("Create", login)
	return nil
}

func (u *LoginRepo) First(db *gorm.DB, userName string) (entities.Login, error) {
	var login entities.Login
	if err := db.Where("user_name = ? ", userName).First(&login).Error; err != nil {
		logger.ErrorLog("An error occured while getting first login - First - loginRepo.go ", "user_name = ? ", userName, err.Error())
		return login, err
	}
	logger.TransactionLog("Create", login)
	return login, nil
}
