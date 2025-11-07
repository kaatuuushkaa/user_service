package userService

import (
	"gorm.io/gorm"
	"user_service/domain"
)

type UserRepository interface {
	GetUserById(id string) (domain.User, error)
	GetLeaderboard() ([]domain.User, error)
	UpdatePoints(userID string, points int) (domain.User, error)
	SetReferrer(userID, referrerID string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetUserById(id string) (domain.User, error) {
	var user domain.User
	err := u.db.First(&user, "id = ?", id).Error
	return user, err
}

func (u *userRepository) GetLeaderboard() ([]domain.User, error) {
	var users []domain.User
	err := u.db.Select("id, name, points, referrer_id").Order("points DESC").Find(&users).Error
	return users, err
}

func (u *userRepository) UpdatePoints(userID string, points int) (domain.User, error) {
	var user domain.User
	err := u.db.Model(&domain.User{}).
		Where("id = ?", userID).
		UpdateColumn("points", gorm.Expr("points + ?", points)).
		Scan(&user).Error
	return user, err
}

func (u *userRepository) SetReferrer(userID, referrerID string) error {
	return u.db.Model(&domain.User{}).
		Where("id = ?", userID).
		Update("referrer_id", referrerID).Error
}
