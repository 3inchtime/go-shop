package service

import (
	"user/internal/model"
	"user/internal/repository"
	"user/internal/utils"
)

type IUserService interface {
	AddUser(user *model.User) (int, error)
	CheckPwd(userName string ,pwd string) (isOk bool,err error)
	GetUser(userID int, userName string) (*model.User, error)
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{UserRepository: userRepository}
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func (u *UserService) AddUser (user *model.User) (int, error) {
	pwdByte, err := utils.GenPassword(user.HashPassword)
	if err != nil {
		return 0, err
	}
	user.HashPassword = string(pwdByte)
	userID, err := u.UserRepository.CreateUser(user)
	return userID, err
}

func (u *UserService) CheckPwd (userName string ,pwd string) (isOk bool,err error){
	user,err:=u.UserRepository.FindUserByName(userName)
	if err!=nil {
		return false,err
	}
	return utils.ValidPassword(pwd,user.HashPassword)
}

func (u *UserService) GetUser (userID int, userName string) (*model.User, error) {
	if userName == "" {
		user, err := u.UserRepository.FindUserByName(userName)
		if err!=nil {
			return nil,err
		}
		return user, nil
	} else {
		user, err := u.UserRepository.FindUserByID(userID)
		if err!=nil {
			return nil,err
		}
		return user, nil
	}
}
