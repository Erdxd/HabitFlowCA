package service

import (
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/repositories"
	"errors"
)

type HabitResponse struct {
	Habit_Name   string
	Status_Today bool
	Streak       int
}

type HabitService struct {
	repoS repositories.Habit
}

func NewHabitService(repo repositories.Habit) *HabitService {
	return &HabitService{repoS: repo}
}
func (s *HabitService) CheckHabit(id_user int) ([]models.HabitFlow, error) {
	return s.repoS.CheckHabit(id_user)
}

func (s *HabitService) AddHabitS(Habit HabitResponse, User_Id int) error {
	if len(Habit.Habit_Name) == 0 {
		return errors.New("Field `name` should be filled ")
	}
	Habits := models.HabitFlow{
		Habit_Name:   Habit.Habit_Name,
		Status_Today: Habit.Status_Today,
		Streak:       Habit.Streak,
		User_Id:      User_Id,
	}

	return s.repoS.AddHabit(Habits, User_Id)
}
func (s *HabitService) DeleteHabitS(id, id_user int) error {
	return s.repoS.DeleteHabit(id, id_user)
}
func (s *HabitService) ChangeStatusHabit(id, streak, id_user int) error {
	return s.repoS.ChangeStatusHabit(id, streak, id_user)
}
func (s *HabitService) GetStreakHabit(user_id, id int) (int, error) {
	return s.repoS.GetStreakHabit(user_id, id)
}
func (s *HabitService) ResetAllStatusHabit() error {
	return s.repoS.ResetAllStatusHabit()
}
func (s *HabitService) GetHabitsByTgId(TgId int64) ([]models.HabitFlow, error) {
	return s.repoS.GetHabitsByTgId(TgId)
}
