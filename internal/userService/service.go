package userService

import (
	"github.com/sirupsen/logrus"
	"user_service/domain"
)

type UserService interface {
	GetUserById(id string) (domain.User, error)
	GetLeaderboard() ([]domain.User, error)
	PostTaskComplete(id string) (domain.User, error)
	PostReferrerHandler(userID, referrerID string) ([]domain.User, error)
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

func (us *userService) PostTaskComplete(id string) (domain.User, error) {
	user, err := us.repo.PostTaskComplete(id)
	if err != nil {
		logrus.Errorf("[users] Failed to complete task: %v", err)
		return domain.User{}, err
	}

	return user, nil
}

func (us *userService) PostReferrerHandler(userID, referrerID string) ([]domain.User, error) {
	_, err := us.repo.GetUserById(userID)
	if err != nil {
		logrus.Errorf("[users] user not found")
		return nil, err
	}
	users, err := us.repo.PostReferrerHandler(userID, referrerID)
	if err != nil {
		logrus.Errorf("[users] Failed to referrer user")
		return nil, err
	}
	return users, nil
}
