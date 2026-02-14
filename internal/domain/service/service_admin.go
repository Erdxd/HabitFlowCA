package service

import (
	"HabitFlow/internal/auth"
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/repositories"
	"errors"
)

type AdminService struct {
	repoA  repositories.Admin
	hasher auth.Hashing
}

func NewAdminService(repo repositories.Admin, hasher auth.Hashing) *AdminService {
	return &AdminService{repoA: repo, hasher: hasher}
}
func (AS *AdminService) GetDataAboutAllUsers() ([]models.User, error) {
	return AS.repoA.GetDataAboutAllUsers()
}
func (AS *AdminService) ChangePasswordForUser(id_user int, newpassword string) error {
	newpasswordhashed, err := AS.hasher.Hash(newpassword)
	if err != nil {
		return errors.New("Wrong password")
	}
	return AS.repoA.ChangePasswordForUser(id_user, newpasswordhashed)
}
func (AS *AdminService) DeleteAccount(id_user int) error {
	return AS.repoA.DeleteAccount(id_user)
}
func (AS *AdminService) DeleteAllHabits(id_user int) error {
	return AS.repoA.DeleteAllHabits(id_user)
}
