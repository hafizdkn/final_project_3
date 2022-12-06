package user

import (
	"final_project_3/auth"
	"final_project_3/helper"
)

type IService interface {
	CreateUser(input UserRegisterInput) (UserResponse, error)
	UserLogin(input UserLoginInput) (UserResponse, error)
	UpdateUser(input UserUpdateInput) (UserResponse, error)
	GetUserById(input int) (User, error)
	GetUsers() ([]User, error)
	DeleteUser(id int) error
}

type service struct {
	repository  IRepository
	genPassword helper.IGenPassword
	genJwt      auth.IService
}

func NewUserService(repository IRepository, jwtService auth.IService) *service {
	genPassword := helper.NewGenPassword
	return &service{repository: repository, genPassword: genPassword, genJwt: jwtService}
}

func (s *service) CreateUser(input UserRegisterInput) (UserResponse, error) {
	var userResponse UserResponse
	var user User

	passwordHashed, err := s.genPassword.GeneratePasswordHash(input.Password)
	if err != nil {
		return userResponse, err
	}

	user = User{
		FullName: input.FullName,
		Email:    input.Email,
		Password: passwordHashed,
		Role:     input.Role,
	}

	user, err = s.repository.CreateUser(user)
	if err != nil {
		return userResponse, err
	}

	userResponse = UserResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: &user.CreatedAt,
	}

	return userResponse, nil
}

func (s *service) UserLogin(input UserLoginInput) (UserResponse, error) {
	var userResponse UserResponse
	var jwtInput auth.JwtInput

	email := input.Email
	InputPassword := input.Password

	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return userResponse, err
	}

	if err := s.genPassword.ComparePasswordHash(user.Password, InputPassword); err != nil {
		return userResponse, err
	}

	jwtInput = auth.JwtInput{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
	}

	token, err := s.genJwt.GenerateToken(jwtInput)
	if err != nil {
		return userResponse, err
	}

	userResponse = UserResponse{
		Token: token,
	}

	return userResponse, nil
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) UpdateUser(input UserUpdateInput) (UserResponse, error) {
	var userResponse UserResponse

	currentUser, err := s.repository.GetUserByEmail(input.EmailCurrentUser)
	if err != nil {
		return userResponse, err
	}

	passwordHashed, err := s.genPassword.GeneratePasswordHash(input.Password)
	if err != nil {
		return userResponse, err
	}

	currentUser.FullName = input.FullName
	currentUser.Password = passwordHashed
	currentUser.Email = input.Email

	updatedUser, err := s.repository.UpdateUser(currentUser)
	if err != nil {
		return userResponse, err
	}

	userResponse = UserResponse{
		ID:        updatedUser.ID,
		FullName:  updatedUser.FullName,
		Email:     updatedUser.Email,
		UpdatedAt: &updatedUser.UpdatedAt,
	}

	return userResponse, nil
}

func (s *service) DeleteUser(id int) error {
	if err := s.repository.Deleteuser(id); err != nil {
		return err
	}

	return nil
}

func (s *service) GetUsers() ([]User, error) {
	users, err := s.repository.GetUsers()
	if err != nil {
		return users, err
	}

	return users, nil
}
