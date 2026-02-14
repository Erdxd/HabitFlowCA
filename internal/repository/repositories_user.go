package repository

import (
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/repositories"
	"database/sql"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.User {
	return &UserRepository{db: db}
}
func (us *UserRepository) Register(Users models.User) error {
	Sqlstatement := (`INSERT INTO "users" (username,email,password) VALUES ($1, $2, $3)`)

	_, err := us.db.Exec(Sqlstatement, Users.Username, Users.Email, Users.Password)
	if err != nil {
		return err
	}
	return nil

}
func (us *UserRepository) Login(Username string) (string, error) {
	var Password string
	Sqlstatement := (`SELECT password FROM users WHERE username = $1`)
	err := us.db.QueryRow(Sqlstatement, Username).Scan(&Password)
	if err != nil {
		return "", err
	}

	return Password, nil
}
func (us *UserRepository) GetUserId(Username string) (int, error) {
	var Id_user int
	Sqlstatement := (`SELECT id_user FROM "users" WHERE username = $1`)
	err := us.db.QueryRow(Sqlstatement, Username).Scan(&Id_user)
	if err != nil {
		return 0, err
	}
	return Id_user, nil
}
func (us *UserRepository) SaveTelegramChatID(userID int, chatID int64) error {
	Sqlstatement := (`UPDATE "users" SET telegram_chat_id = $1 WHERE id_user = $2`)
	_, err := us.db.Exec(Sqlstatement, chatID, userID)
	return err
}
func (us *UserRepository) GetTelegramChatID(userID int) (int64, error) {
	var TgChat int
	Sqlstatement := (`SELECT telegram_chat_id FROM "users" WHERE id_user = $1`)
	err := us.db.QueryRow(Sqlstatement, userID).Scan(&TgChat)
	if err != nil {
		return 0, err
	}
	return int64(TgChat), nil

}
func (us *UserRepository) GetUserIdByTgID(chatId int64) (int, error) {

	var User_Id int
	Sqlstatement := (`SELECT id_user FROM "users" WHERE telegram_chat_id = $1`)
	err := us.db.QueryRow(Sqlstatement, chatId).Scan(&User_Id)
	if err != nil {
		return 0, err
	}
	return User_Id, nil

}
func (us *UserRepository) GetDataUser(id_user int) ([]models.UserBaseView, error) {
	log.Println("GetDataUser called with id_user:", id_user)
	rows, err := us.db.Query(`SELECT id_user,username,email,password FROM users WHERE id_user = $1`, id_user)

	if err != nil {
		log.Println("Can't SELECT data by your tables")
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var BaseUser []models.UserBaseView
	for rows.Next() {
		var U models.UserBaseView
		err := rows.Scan(&U.Id_user, &U.Username, &U.Email, &U.Password)
		if err != nil {
			return nil, err
		}
		BaseUser = append(BaseUser, U)

	}

	return BaseUser, nil
}
func (us *UserRepository) RedactLogin(user_id int, username string) error {
	SqlStatement := (`UPDATE "users" SET username = $1 WHERE id_user = $2`)
	_, err := us.db.Exec(SqlStatement, username, user_id)
	if err != nil {
		return err
	}
	return nil

}
func (us *UserRepository) GetPasswordwithId(Id_user int) (string, error) {
	var Password string
	Sqlstatement := (`SELECT password FROM "users" WHERE id_user =$1`)
	err := us.db.QueryRow(Sqlstatement, Id_user).Scan(&Password)
	if err != nil {
		return "", err
	}

	return Password, nil

}
func (us *UserRepository) RedactPassword(NewHashedPassword string, id_user int) error {
	SqlStatement := (`UPDATE users SET password = $1 WHERE id_user = $2`)
	_, err := us.db.Exec(SqlStatement, NewHashedPassword, id_user)
	if err != nil {

		return err
	}
	return nil
}
func (us *UserRepository) GetAdmin(id_user int) (bool, error) {
	var Admin bool
	SqlStatement := (`SELECT admin from users WHERE id_user = $1`)
	err := us.db.QueryRow(SqlStatement, id_user).Scan(&Admin)
	if err != nil {
		return false, err
	}
	return Admin, nil
}
