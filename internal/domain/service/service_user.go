package service

import (
	"HabitFlow/internal/auth"
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/repositories"
	"errors"
	"log"
)

type UserService struct {
	usersRepo repositories.User
	hasher    auth.Hashing
	jwt       auth.JWT
}
type UserRegisterResponse struct {
	Id_user  int
	Username string
	Email    string
	Password string
}

func NewuserService(users repositories.User, hasher auth.Hashing) *UserService {
	return &UserService{usersRepo: users, hasher: hasher}
}
func (usS *UserService) Register(userh UserRegisterResponse, AgainPassword string) error {

	if userh.Password != AgainPassword {
		return errors.New("Not matched with original password")

	}
	hashed, err := usS.hasher.Hash(userh.Password)
	if err != nil {
		log.Println(err)
		return err

	}

	userh.Password = hashed
	user := models.User{
		Username: userh.Username,
		Email:    userh.Email,
		Password: userh.Password,
	}
	return usS.usersRepo.Register(user)

}
func (usS *UserService) Login(Username string, passwordfromuser string) error {
	passwordfromdb, err := usS.usersRepo.Login(Username)

	if err != nil {
		return errors.New("Wrong Login or Password")
	}
	compare := usS.hasher.Compare(passwordfromdb, passwordfromuser)
	if !compare {
		return errors.New("Wrong Login or Password")
	}
	return nil

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
func (usS *UserService) RedactLogin(user_id int, username, Passworduser string) error {
	passwordFromdb, err := usS.usersRepo.GetPasswordwithId(user_id)
	if err != nil {
		return errors.New("Something wrong")
	}
	Coincidence := usS.hasher.Compare(passwordFromdb, Passworduser)

	if !Coincidence {
		return errors.New("Worng Password")
	}

	return usS.usersRepo.RedactLogin(user_id, username)
}

func (usS *UserService) GetPasswordwithId(Id_user int) (string, error) {
	return usS.usersRepo.GetPasswordwithId(Id_user)
}
func (usS *UserService) RedactPassword(Newpassword string, id_user int, Oldpassword string) error {
	PasswordFromdb, err := usS.usersRepo.GetPasswordwithId(id_user)
	if err != nil {
		return errors.New("Something wrong")
	}
	Coincidence := usS.hasher.Compare(PasswordFromdb, Oldpassword)
	if !Coincidence {
		return errors.New("Wrong Password")

	}
	NewpasswordHash, err := usS.hasher.Hash(Newpassword)
	if err != nil {
		return errors.New("Something wrong")
	}
	return usS.usersRepo.RedactPassword(NewpasswordHash, id_user)
}
func (usS *UserService) GetAdmin(id_user int) (bool, error) {

	return usS.usersRepo.GetAdmin(id_user)
}
