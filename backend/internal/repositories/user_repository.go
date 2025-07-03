package repositories

import (
	"github.com/jadzon/burger_shop/internal/database"
	"github.com/jadzon/burger_shop/internal/models"
)

type userRepository struct {
	db *database.Database
}
type UserRepository interface {
	Create(user *models.User) (*models.User, error)
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepository{
		db: &db,
	}
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	tx := r.db.Create(user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}
