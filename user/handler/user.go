package handler

import (
	"context"
	"user/internal/model"
	"user/internal/service"

	pb "user/proto"
)

type User struct {
	UserService service.IUserService
}

func (u *User) Register(ctx context.Context, userRegisterRequest *pb.UserRegisterRequest, userRegisterResponse *pb.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     userRegisterRequest.UserName,
		Email:        userRegisterRequest.Email,
		Telephone:    userRegisterRequest.Telephone,
		Age:          int(userRegisterRequest.Age),
		HashPassword: userRegisterRequest.Pwd,
	}
	_, err := u.UserService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userRegisterResponse.Message = "Create User Success"
	return nil
}

func (u *User) Login(ctx context.Context, userLoginRequest *pb.UserLoginRequest, userLoginResponse *pb.UserLoginResponse) error {
	isOk, err := u.UserService.CheckPwd(userLoginRequest.UserName, userLoginRequest.Pwd)
	if err != nil {
		return err
	}
	userLoginResponse.IsSuccess = isOk
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *pb.UserInfoRequest, userInfoResponse *pb.UserInfoResponse) error {
	user, err := u.UserService.GetUser(int(userInfoRequest.UserId), userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse = userForResponse(user)
	return nil
}

func userForResponse(user *model.User) *pb.UserInfoResponse {
	response := &pb.UserInfoResponse{}
	response.UserName = user.UserName
	response.UserId = int64(user.ID)
	response.Email = user.Email
	response.Telephone = user.Telephone
	response.Age = int32(user.Age)
	return response
}
