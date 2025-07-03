package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Code string `json:"code"`
}
type RegisterUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type UserResponse struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}
