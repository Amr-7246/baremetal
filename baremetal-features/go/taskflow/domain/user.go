package domain

import "time"

//& The full entity
type User struct {
	ID             string    `json:"id"` //? why does json tags here + tell me more about that syntax
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"` // Never send to client
	FullName       string    `json:"full_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
//& for creating new users
type UserCreateDTO struct { //? what does DTO refer to in userDTO
	Email     string `json:"email"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
}
//& UserResponseDTO for API responses (excludes sensitive data)
type UserResponseDTO struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}
//& ToResponse converts User to UserResponseDTO
func (u *User) ToResponse() *UserResponseDTO {
	return &UserResponseDTO{
		ID:        u.ID,
		Email:     u.Email,
		FullName:  u.FullName,
		CreatedAt: u.CreatedAt,
	}
}