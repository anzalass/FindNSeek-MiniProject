package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `json:"id" form:"id`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserInterface interface {
	Register(data User) (*User, error)
	Login(data User) (*User, error)
}

type UserModel struct {
	db *gorm.DB
}

func (um *UserModel) Init(db *gorm.DB) {
	um.db = db
}

func NewUserModel(db *gorm.DB) UserInterface {
	return &UserModel{
		db: db,
	}
}

func (um *UserModel) Register(data User) (*User, error) {
	if err := um.db.Create(&data).Error; err != nil {
		logrus.Error("model : error register user")
		return nil, err
	}

	return &data, nil
}

func (um *UserModel) Login(data User) (*User, error) {

	if err := um.db.Where("email = ? AND password = ?", data.Email, data.Password).First(&data).Error; err != nil {
		logrus.Error("model : error login user")
		return nil, err
	}

	return &data, nil
}
