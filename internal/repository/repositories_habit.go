package repository

import (
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/repositories"
	"database/sql"
	"log"
)

type HabitRepository struct {
	db *sql.DB
}

func NewHabitRepository(db *sql.DB) repositories.Habit {
	return &HabitRepository{db: db}
}
func (r *HabitRepository) CheckHabit(Id_user int) ([]models.HabitFlow, error) {
	rows, err := r.db.Query(`SELECT id, habit_name, status_today, streak FROM "HabitFlow" WHERE user_id= $1`, Id_user)

	if err != nil {
		log.Println("Can't SELECT data by your tables")
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var habits []models.HabitFlow
	for rows.Next() {
		var habit models.HabitFlow
		err := rows.Scan(&habit.Id, &habit.Habit_Name, &habit.Status_Today, &habit.Streak)
		if err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}
	return habits, nil

}
func (r *HabitRepository) AddHabit(Habits models.HabitFlow, Id_user int) error {
	SqlStatement := (`INSERT INTO "HabitFlow" (id, habit_name, status_today, streak, user_id) VALUES ($1,$2 ,$3,$4, $5)`)
	_, err := r.db.Exec(SqlStatement, Habits.Id, Habits.Habit_Name, Habits.Status_Today, Habits.Streak, Id_user)
	if err != nil {
		return err
	}
	return nil
}
func (r *HabitRepository) DeleteHabit(id, user_id int) error {
	SqlStatement := (`DELETE FROM "HabitFlow" WHERE id = $1 AND user_id = $2 `)
	_, err := r.db.Exec(SqlStatement, id, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (r *HabitRepository) ChangeStatusHabit(id, streak, user_id int) error {
	SqlStatement := (`UPDATE "HabitFlow" SET status_today = true, streak = $1 WHERE id = $2 AND user_id = $3`)
	_, err := r.db.Exec(SqlStatement, streak+1, id, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (r *HabitRepository) GetStreakHabit(user_id, id int) (int, error) {
	var streak int
	SqlStatement := (`SELECT streak FROM "HabitFlow" WHERE id = $1 AND user_id = $2`)
	err := r.db.QueryRow(SqlStatement, id, user_id).Scan(&streak)
	if err != nil {
		return 0, nil
	}
	return streak, nil
}
func (r *HabitRepository) ResetAllStatusHabit() error {
	SqlStatement := (`UPDATE "HabitFlow" SET status_today = false`)
	_, err := r.db.Exec(SqlStatement)
	if err != nil {
		return err
	}
	return nil

}
func (r *HabitRepository) DeleteAllHabits(user_id int) error {
	SqlStatement := (`DELETE FROM "HabitFlow" WHERE user_id = $1`)
	_, err := r.db.Exec(SqlStatement, user_id)
	if err != nil {
		return err
	}
	return nil

}

