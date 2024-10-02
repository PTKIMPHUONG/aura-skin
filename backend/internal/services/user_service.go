package services

import (
	"auraskin/internal/models"
	"auraskin/internal/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *models.User) error
	Login(email string) (*models.User, error)
	Logout(userID string) error
	DeleteAccount(userID string) error
	Update(userID string, user *models.User) error
	ForgotPassword(email string) error
	ChangePassword(userID, oldPassword, newPassword string) error
}

type userService struct {
	repo repositories.UserRepository 
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

// Register a new user 
func (s *userService) Register(user *models.User) error {
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	// Pass the user to the repository to create a new node
	return s.repo.CreateUserNode(*user)
}

// Login a user by email 
func (s *userService) Login(email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) Logout(userID string) error {
	return nil
}

// Delete a user account 
func (s *userService) DeleteAccount(userID string) error {
	return s.repo.DeleteUserNode(userID)
}

// Update a user's information 
func (s *userService) Update(userID string, user *models.User) error {
	return s.repo.UpdateUserNode(userID, *user)
}

// ForgotPassword 
func (s *userService) ForgotPassword(email string) error {
	return nil
}

// Change a user's password 
func (s *userService) ChangePassword(userID, oldPassword, newPassword string) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Check if the old password matches
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("invalid old password")
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password in the database
	user.Password = string(hashedPassword)
	return s.repo.UpdateUserNode(userID, user)
}
