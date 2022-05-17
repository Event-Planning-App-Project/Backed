package user

import (
	"errors"
	"event/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (ur *UserRepo) Login(email string, password string) (entities.User, error) {
	users := []entities.User{}

	if err := ur.Db.Where("email = ? AND password = ?", email, password).First(&users).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("tidak bisa select data")
	}

	return users[0], nil
}

func (ur *UserRepo) InsertUser(newUser entities.User) (entities.User, error) {
	if err := ur.Db.Create(&newUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("tidak bisa insert data")
	}

	log.Info()
	return newUser, nil
}

func (ur *UserRepo) GetUserID(ID int) (entities.User, error) {
	arrUser := []entities.User{}

	if err := ur.Db.Where("id = ?", ID).Find(&arrUser).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("tidak bisa select data")
	}

	if len(arrUser) == 0 {
		log.Warn("data tidak ditemukan")
		return entities.User{}, errors.New("data tidak ditemukan")
	}

	log.Info()
	return arrUser[0], nil
}

func (ur *UserRepo) UpdateUser(ID int, update entities.User) (entities.User, error) {
	var res entities.User
	if err := ur.Db.Where("id = ?", ID).Updates(&update).Find(&res).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("tidak bisa update data")
	}

	log.Info()
	return res, nil
}

func (pr *UserRepo) DeleteUser(ID int) (entities.User, error) {
	var user []entities.User
	res, err := pr.GetUserID(ID)
	if err != nil {
		return entities.User{}, err
	}

	if err := pr.Db.Delete(&user, "id = ?", ID).Error; err != nil {
		log.Warn(err)
		return entities.User{}, errors.New("tidak bisa delete data")
	}
	log.Info()
	return res, nil

}
