package userService

import (
	"github.com/sirupsen/logrus"
	"user_service/domain"
)

type UserService interface {
	GetUserById(id string) (domain.User, error) //GET /users/{id}/status
	GetLeaderboard() ([]domain.User, error)     //GET /users/leaderboard
	//PostTask                                    //POST /users/{id}/task/complete
	//PostReferrer                                //POST /users/{id}/referrer
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (us *userService) GetUserById(id string) (domain.User, error) {
	return us.repo.GetUserById(id)
}

func (us *userService) GetLeaderboard() ([]domain.User, error) {
	users, err := us.repo.GetLeaderboard()
	if err != nil {
		logrus.Errorf("[users] Failed to fetch users from DB: %v", err)
		return nil, err
	}

	return users, nil
}
