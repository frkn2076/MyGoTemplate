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
		return err
	}
	return nil
}

func (u *LoginRepo) First(db *gorm.DB, userName string) (entities.Login, error) {
	var login entities.Login
	if err := db.Where("user_name = ? ", userName).First(&login).Error; err != nil {
		return login, err
	} else {
		return login, nil
	}
}
