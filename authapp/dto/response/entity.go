package response

import (
	"authapp/domain/user"
	"time"
)

type (
	UserResponse struct {
		ID          int    `json:"id"`
		Phonenumber string `json:"phonenumber"`
		Name        string `json:"name"`
		Role        string `json:"role"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
	}
)

func FromUserToResponse(u *user.User) *UserResponse {
	return &UserResponse{
		ID:          u.ID,
		Phonenumber: u.Phonenumber,
		Name:        u.Name,
		Role:        string(u.Role),
		CreatedAt:   u.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   u.UpdatedAt.Format(time.RFC3339),
	}
}
