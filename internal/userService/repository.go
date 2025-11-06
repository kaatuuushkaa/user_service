package userService

import (
	"gorm.io/gorm"
	"user_service/domain"
)

type UserRepository interface {
	GetUserById(id string) (domain.User, error) //GET /users/{id}/status
	GetLeaderboard() ([]domain.User, error)     //GET /users/leaderboard
	//PostTask                                    //POST /users/{id}/task/complete
	//PostReferrer                                //POST /users/{id}/referrer
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepositiry(db *gorm.DB) UserRepository {
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
