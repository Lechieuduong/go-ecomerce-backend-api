package service

import "go-ecomerce-backend-api/m/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUserByID() string {
	return us.userRepo.GetInfoUserByID()
}