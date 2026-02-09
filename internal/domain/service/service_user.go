package service

import (
	"HabitFlow/internal/auth"
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/repositories"
	"log"
)

type UserService struct {
	usersRepo repositories.User
	hasher    auth.Hashing
}

func NewuserService(users repositories.User, hasher auth.Hashing) *UserService {
	return &UserService{usersRepo: users, hasher: hasher}
}
func (usS *UserService) Register(user models.User) error {
	hashed, err := usS.hasher.Hash(user.Password)
	if err != nil {
		log.Println(err)
		return err

	}
	user.Password = hashed
	return usS.usersRepo.Register(user)

}
func (usS *UserService) GetPassword(Username string) (string, error) {
	return usS.usersRepo.GetPassword(Username)
}
func (usS *UserService) GetUserId(Username string) (int, error) {
	return usS.usersRepo.GetUserId(Username)
}
func (usS *UserService) SaveTelegramChatID(userId int, chatId int64) error {
	return usS.usersRepo.SaveTelegramChatID(userId, chatId)
}
func (usS *UserService) GetTelegramChatID(user_id int) (int64, error) {
	return usS.usersRepo.GetTelegramChatID(user_id)

}
func (usS *UserService) GetUserIdByTgID(chatId int64) (int, error) {
	return usS.usersRepo.GetUserIdByTgID(chatId)
}
func (usS *UserService) GetDataUser(id_user int) ([]models.UserBaseView, error) {
	return usS.usersRepo.GetDataUser(id_user)
}
func (usS *UserService) RedactLogin(user_id int, username string) error {
	return usS.usersRepo.RedactLogin(user_id, username)
}

func (usS *UserService) GetPasswordwithId(Id_user int) (string, error) {
	return usS.usersRepo.GetPasswordwithId(Id_user)
}
func (usS *UserService) RedactPassword(NewHashedPassword string, id_user int) error {
	return usS.usersRepo.RedactPassword(NewHashedPassword, id_user)
}
