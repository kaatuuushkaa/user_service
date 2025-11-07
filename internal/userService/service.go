package userService

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
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
	user, err := us.repo.UpdatePoints(id, 15)
	if err != nil {
		logrus.Errorf("[users] Failed to complete task: %v", err)
		return domain.User{}, err
	}

	return user, nil
}

func (us *userService) PostReferrerHandler(userID, referrerID string) ([]domain.User, error) {
	user, err := us.repo.GetUserById(userID)
	if err != nil {
		logrus.Errorf("[users] user not found")
		return nil, err
	}
	if user.ReferrerID > 0 {
		logrus.Infof("[users] user %s already has referrer, skipping", userID)
		referrer, _ := us.repo.GetUserById(strconv.Itoa(user.ReferrerID))
		return []domain.User{user, referrer}, nil
	}
	referrer, err := us.repo.GetUserById(referrerID)
	if err != nil {
		logrus.Warnf("[users] referrer %s not found", referrerID)
		return nil, fmt.Errorf("referrer not found")
	}

	if referrer.ReferrerID > 0 && referrer.ReferrerID == user.ID {
		return nil, fmt.Errorf("cyclic referral is not allowed: user %s is already referrer of %s", referrerID, userID)
	}

	if user.ID == referrer.ID {
		return nil, fmt.Errorf("user cannot refer themselves")
	}

	if err := us.repo.SetReferrer(userID, referrerID); err != nil {
		return nil, err
	}

	updatedUser, err := us.repo.UpdatePoints(userID, 5)
	if err != nil {
		return nil, err
	}

	updatedReferrer, err := us.repo.UpdatePoints(referrerID, 10)
	if err != nil {
		return nil, err
	}

	return []domain.User{updatedUser, updatedReferrer}, nil
}
