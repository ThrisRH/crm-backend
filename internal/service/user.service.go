package service

import "go-sea-crm/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserInfoService() string {
	return us.userRepo.GetUserInfoRepo()
}	

