package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(id int, file_location string) (User, error)
	GetUserById(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password)
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Email == "" {
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	user.Password = string(password)
	newUser, err := s.repository.FindByEmail(user.Email)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Email == "" {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(id int, file_location string) (User, error) {

	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Avatar = file_location

	updated_user, err := s.repository.Update(user)
	if err != nil {
		return updated_user, err
	}

	return updated_user, nil
}

func (s *service) GetUserById(id int) (User, error) {

	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.Email == "" {
		return user, errors.New("no user found")
	}

	return user, nil
}
