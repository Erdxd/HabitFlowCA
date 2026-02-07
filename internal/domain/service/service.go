package service

import (
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/repositories"
	"HabitFlow/internal/repository"
	"errors"
)

type HabitService struct {
	repoS repositories.Habit
}

func NewHabitService(repo *repository.HabitRepository) *HabitService {
	return &HabitService{repoS: repo}
}
func (s *HabitService) CheckHabit(id_user int) ([]models.HabitFlow, error) {
	return s.repoS.CheckHabit(id_user)
}

func (s *HabitService) AddHabitS(Habit models.HabitFlow, User_Id int) error {
	if len(Habit.Habit_Name) == 0 {
		return errors.New("Field `name` should be filled ")
	}
	return s.repoS.AddHabit(Habit, User_Id)
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
func (s *HabitService) DeleteAllHabits(user_id int) error {
	return s.repoS.DeleteAllHabits(user_id)
}
