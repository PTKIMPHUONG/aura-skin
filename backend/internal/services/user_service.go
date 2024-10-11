package services

import (
	"auraskin/internal/models"
	"auraskin/internal/repositories"
	"errors"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *models.User) error
	Login(email string) (*models.User, error)
	Logout(userID string) error
	DeleteAccount(userID string) error
	Update(user *models.User) error
	ForgotPassword(email string) error
	ChangePassword(userID, oldPassword, newPassword string) error
	ComparePassword(hashedPassword string, plainPassword string) error
	GetOrdersByUserID(id string) ([]models.Order, error)
	UploadProfilePicture(userID string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) // Thêm chức năng upload ảnh đại diện

}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}


// Register a new user
func (s *userService) Register(user *models.User) error {
	// Kiểm tra nếu các trường cần thiết có dữ liệu hợp lệ
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if user.PhoneNumber == "" {
		return errors.New("phonenumber is required")
	}

	// Mã hóa mật khẩu trước khi lưu vào database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// Gọi repository để tạo người dùng mới
	err = s.repo.Create(*user)
	if err != nil {
		return err
	}

	return nil
}

// Login a user by email
func (s *userService) Login(email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) ComparePassword(hashedPassword string, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

func (s *userService) Logout(userID string) error {
	return nil
}

// Delete a user account
func (s *userService) DeleteAccount(userID string) error {
	existingUser, err := s.repo.GetByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if userID != existingUser.ID {
		return errors.New("user ID does not match")
	}
	return s.repo.Delete(userID)
}

// Update a user's information
func (s *userService) Update(user *models.User) error {
	existingUser, err := s.repo.GetByID(user.ID)
	if err != nil {
		return errors.New("user not found")
	}

	if user.ID != existingUser.ID {
		return errors.New("user ID does not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	if err := s.repo.Update(*user); err != nil {
		return errors.New("failed to update user information")
	}

	return nil
}

// ForgotPassword
func (s *userService) ForgotPassword(email string) error {
	return nil
}

// Change a user's password
func (s *userService) ChangePassword(userID, oldPassword, newPassword string) error {
	user, err := s.repo.GetByID(userID)
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
	return s.repo.Update(*user)
}

func (s *userService) GetOrdersByUserID(id string) ([]models.Order, error) {
	return s.repo.GetOrdersByUserID(id)
}
func (s *userService) UploadProfilePicture(userID string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	return s.repo.UploadProfilePicture(userID, file, fileHeader)
}