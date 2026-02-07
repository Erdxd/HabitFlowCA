package handlers

import (
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/http/middleware"
	"HabitFlow/internal/http/service"
	"log"
	"net/http"
)

type HabitHandler struct {
	HabitService *service.HabitService
}

func NewHabitHandler(HabitService *service.HabitService) *HabitHandler {
	return &HabitHandler{HabitService: HabitService}
}
func (h *HabitHandler) CheckHabit(w http.ResponseWriter, r *http.Request) {
	Id_user, err := middleware.Cookie_userID(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	Habits, err := h.HabitService.CheckHabit(Id_user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		HabitAll []models.HabitFlow
	}{
		HabitAll: Habits,
	}
	tmplmain.Execute(w, data)
}
