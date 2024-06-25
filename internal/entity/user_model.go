package entity

type User struct {
	BaseModel
	Email    string
	Password string
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:    u.ID.String(),
		Email: u.Email,
	}
}
