package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	segmentation_service "segmentation-service"
	"segmentation-service/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAllUsers() ([]segmentation_service.UserResponse, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) CreateUser(user segmentation_service.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("salt"))))
}