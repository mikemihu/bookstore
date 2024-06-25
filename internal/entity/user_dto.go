package entity

type UserFilter struct {
	User
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
