package service

import (
	"example/demo-server/domain/model"
	"example/demo-server/domain/repository"
)

type IUserService interface {
	AddUser(user *model.User) (uint, error)
}

// 创建
func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository}
}

type UserService struct {
	UserRepository repository.IUserRepository
}

// 插入
func (u *UserService) AddUser(user *model.User) (uint, error) {
	return u.UserRepository.CreateUser(user)
}
