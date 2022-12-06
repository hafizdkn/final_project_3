package user

type UserRegisterInput struct {
	FullName string `json:"full_name" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateInput struct {
	FullName         string `json:"full_name" binding:"required,min=5"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=6"`
	EmailCurrentUser string
}
