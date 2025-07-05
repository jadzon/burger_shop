package services

import (
	"github.com/jadzon/burger_shop/internal/dto"
	"github.com/jadzon/burger_shop/internal/models"
	"github.com/jadzon/burger_shop/internal/repositories"
	"github.com/jadzon/burger_shop/internal/security"
)

type userService struct {
	userRepo repositories.UserRepository
}
type UserService interface {
	Create(req *dto.CreateUserRequest) (*dto.UserResponse, error)
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
func (s *userService) Create(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	userModel := mapCreateUserRequestToModel(req)
	codeHash, err := security.HashCode(req.Code)
	if err != nil {
		return nil, err
	}
	userModel.CodeHash = codeHash
	user, err := s.userRepo.Create(userModel)
	if err != nil {
		return nil, err
	}
	userResponseDTO := mapUserModelToResponseDTO(*user)
	return userResponseDTO, nil
}
func mapCreateUserRequestToModel(req *dto.CreateUserRequest) *models.User {
	return &models.User{
		Name: req.Name,
		Age:  req.Age,
	}
}
func mapUserModelToResponseDTO(model models.User) *dto.UserResponse {
	return &dto.UserResponse{
		Id:        model.Id,
		Name:      model.Name,
		Age:       model.Age,
		CreatedAt: model.CreatedAt,
	}
}
