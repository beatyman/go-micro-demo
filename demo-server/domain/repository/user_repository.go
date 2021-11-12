package repository

import (
	"example/demo-server/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InitTable() error
	CreateUser(user *model.User) (uint, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDb.AutoMigrate(&model.User{})
}

// 创建Category 信息
func (u *UserRepository) CreateUser(category *model.User) (uint, error) {
	return category.ID, u.mysqlDb.Create(category).Error
}
