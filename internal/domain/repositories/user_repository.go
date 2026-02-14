package repositories

import (
	"HabitFlow/internal/domain/models"
)

type User interface {
	Register(user models.User) error
	Login(Username string) (string, error)
	GetUserId(Username string) (int, error)
	SaveTelegramChatID(userID int, chatID int64) error
	GetTelegramChatID(userID int64) (int64, error)
	GetUserIdByTgID(chatId int64) (int, error)
	GetDataUser(id_user int) ([]models.UserBaseView, error)
	GetAdmin(id_user int) (bool, error)
	RedactLogin(user_id int, username string) error
	GetPasswordwithId(Id_user int) (string, error)
	RedactPassword(NewHashedPassword string, id_user int) error
}
