package repository

import (
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/repositories"
	"database/sql"
)

type AdminRepository struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) repositories.Admin {
	return &AdminRepository{db: db}
}

func (A *AdminRepository) GetDataAboutAllUsers() ([]models.User, error) {
	SqlStatement := (`SELECT * FROM users`)
	rows, err := A.db.Query(SqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id_user, &u.Username, &u.Email, &u.Password, &u.Telegram_chat_id, &u.Admin)
		if err != nil {
			return nil, err
		}
		users = append(users, u)

	}
	return users, nil

}

func (A *AdminRepository) ChangePasswordForUser(id_user int, newpasswordhashed string) error {
	SqlStatement := (`UPDATE users SET password = $1 WHERE id_user = $2`)

	_, err := A.db.Exec(SqlStatement, newpasswordhashed, id_user)
	if err != nil {

		return err
	}
	return nil
}
func (A *AdminRepository) DeleteAccount(id_user int) error {
	SqlStatement := (`DELETE FROM users WHERE id_user = $1 `)
	_, err := A.db.Exec(SqlStatement, id_user)
	if err != nil {

		return err
	}
	return nil
}
func (A *AdminRepository) DeleteAllHabits(user_id int) error {
	SqlStatement := (`DELETE FROM "HabitFlow" WHERE user_id = $1`)
	_, err := A.db.Exec(SqlStatement, user_id)
	if err != nil {
		return err
	}
	return nil
}
