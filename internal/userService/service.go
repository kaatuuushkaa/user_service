package userService

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"user_service/domain"
)

type UserService interface {
	GetUserById(id string) (domain.User, error)
	GetLeaderboard() ([]domain.User, error)
	PostTaskComplete(id string) (domain.User, error)
	PostReferrerHandler(userID, referrerID string) ([]domain.User, error)
	SignIn(username, password string) (domain.User, error)
	SignUp(username, password string) (domain.User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (us *userService) SignUp(username, password string) (domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		Username:        username,
		Hashed_password: string(hashedPassword),
	}

	user, err = us.repo.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (us *userService) GetUserById(id string) (domain.User, error) {
	return us.repo.GetUserById(id)
}

func (us *userService) GetLeaderboard() ([]domain.User, error) {
	users, err := us.repo.GetLeaderboard()
	if err != nil {
		logrus.Errorf("Failed to fetch users from DB: %v", err)
		return nil, err
	}

	return users, nil
}

func (us *userService) PostTaskComplete(id string) (domain.User, error) {
	user, err := us.repo.UpdatePoints(id, 15)
	if err != nil {
		logrus.Errorf("Failed to complete task: %v", err)
		return domain.User{}, err
	}

	return user, nil
}

func (us *userService) PostReferrerHandler(userID, referrerID string) ([]domain.User, error) {
	user, err := us.repo.GetUserById(userID)
	if err != nil {
		logrus.Errorf("User not found")
		return nil, err
	}
	if user.ReferrerID != nil {
		return nil, fmt.Errorf("User %s already has referrer", userID)
	}
	referrer, err := us.repo.GetUserById(referrerID)
	if err != nil {
		logrus.Warnf("Referrer %s not found", referrerID)
		return nil, fmt.Errorf("referrer not found")
	}

	if referrer.ReferrerID != nil && *referrer.ReferrerID == user.ID {
		return nil, fmt.Errorf("Cyclic referral is not allowed: user %s is already referrer of %s", referrerID, userID)
	}

	if user.ID == referrer.ID {
		return nil, fmt.Errorf("User cannot refer themselves")
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

func (us *userService) SignIn(username, password string) (domain.User, error) {
	user, err := us.repo.GetUserByUsername(username)
	if err != nil {
		return domain.User{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Hashed_password), []byte(password)); err != nil {
		fmt.Println(password, user.Hashed_password)
		return domain.User{}, err
	}

	return user, nil
}
