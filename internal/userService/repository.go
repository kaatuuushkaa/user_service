package userService

import (
	"gorm.io/gorm"
	"user_service/domain"
)

type UserRepository interface {
	GetUserById(id string) (domain.User, error)
	GetLeaderboard() ([]domain.User, error)
	PostTaskComplete(id string) (domain.User, error)
	PostReferrerHandler(userID, referrerID string) ([]domain.User, error)
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
	err := u.db.Select("id, name, points").Order("points DESC").Find(&users).Error
	return users, err
}

func (u *userRepository) PostTaskComplete(id string) (domain.User, error) {
	var user domain.User
	err := u.db.Model(&user).Where("id=?", id).UpdateColumn("points", gorm.Expr("points+?", 10)).Scan(&user).Error
	return user, err
}

func (u *userRepository) PostReferrerHandler(userID, referrerID string) ([]domain.User, error) {
	var user domain.User
	if err := u.db.Model(&user).Where("id=?", userID).UpdateColumn("points", gorm.Expr("points+?", 5)).Scan(&user).Error; err != nil {
		return nil, err
	}
	var user2 domain.User
	if err := u.db.Model(&user2).Where("id=?", referrerID).UpdateColumn("points", gorm.Expr("points+?", 10)).Error; err != nil {
		return nil, err
	}
	if err := u.db.Model(&user).Where("id=?", userID).UpdateColumn("referrer_id", referrerID).Scan(&user2).Error; err != nil {
		return nil, err
	}
	return []domain.User{user, user2}, nil
}
