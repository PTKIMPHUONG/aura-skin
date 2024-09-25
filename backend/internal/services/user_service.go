package services

// import (
// 	"firstproject/models"
// 	"firstproject/repositories"

// 	"golang.org/x/crypto/bcrypt"
// )

// type UserService interface {
// 	Register(user *models.User) error
// 	FindUserByEmail(email string) (*models.User, error)
// 	DeleteUser(userID string) error
// }

// type userService struct {
// 	repo repositories.UserRepository
// }

// func NewUserService(repo repositories.UserRepository) UserService {
// 	return &userService{repo}
// }

// func (s *userService) Register(user *models.User) error {
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	user.Password = string(hashedPassword)
// 	return s.repo.CreateUser(user)
// }

// func (s *userService) FindUserByEmail(email string) (*models.User, error) {
// 	return s.repo.FindUserByEmail(email)
// }

// func (s *userService) DeleteUser(email string) error {
// 	return s.repo.DeleteUser(email)
// }
