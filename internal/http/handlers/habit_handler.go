package handlers

import (
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/service"
	"HabitFlow/internal/http/middleware"

	"log"
	"net/http"
)

type HabitHandler struct {
	HabitService *service.HabitService
	Auth         *middleware.JWTToken
}

func NewHabitHandler(HabitService *service.HabitService, Auth *middleware.JWTToken) *HabitHandler {
	return &HabitHandler{HabitService: HabitService, Auth: Auth}
}
func (h *HabitHandler) CheckHabit(w http.ResponseWriter, r *http.Request) {
	Id_user, err := h.Auth.Cookie_userID(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
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
