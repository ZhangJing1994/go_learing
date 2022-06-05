package main

import (
	"lession2/week4/internal/biz"
	"lession2/week4/internal/data"
)

// Injectors from wire.go:

func InitUserUsecase() *biz.UserUsecase {
	userRepo := data.NewUserRepo()
	userUsecase := biz.NewUserUsecase(userRepo)
	return userUsecase
}
