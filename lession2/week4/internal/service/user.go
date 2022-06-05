// service/user.go

package service

import (
	"context"
	v1 "lession2/week4/api/user/v1"
	"lession2/week4/internal/biz"
)

type UserService struct {
	u *biz.UserUsecase
	v1.UnimplementedUserServer
}

func NewUserService(u *biz.UserUsecase) v1.UserServer {
	return &UserService{u: u}
}

func (s *UserService) RegisterUser(ctx context.Context, r *v1.AddUserRequest) (*v1.AddUserResponse, error) {
	// dto -> do
	u := &biz.User{Name: r.Name, Age: r.Age}

	// call biz
	s.u.SaveUser(u)

	// return reply
	return &v1.AddUserResponse{Id: u.ID}, nil
}
