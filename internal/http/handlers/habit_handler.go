package handlers

import (
	"HabitFlow/internal/domain/service"
	"HabitFlow/internal/http/dto"
	"HabitFlow/internal/http/middleware"
	"strconv"
	"text/template"

	"log"
	"net/http"
)

type HabitHandler struct {
	HabitService *service.HabitService
	Auth         *middleware.JWTToken
	tmplmain     *template.Template
}

func NewHabitHandler(HabitService *service.HabitService, Auth *middleware.JWTToken, tmpl *template.Template) *HabitHandler {
	return &HabitHandler{HabitService: HabitService, Auth: Auth, tmplmain: tmpl}
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
	var NewHabitDto []dto.ResponseHabitCheck
	for _, habit := range Habits {
		Habits := dto.ResponseHabitCheck{
			Habit_Name:   habit.Habit_Name,
			Status_Today: habit.Status_Today,
			Streak:       habit.Streak,
		}
		NewHabitDto = append(NewHabitDto, Habits)
	}

	data := map[string]interface{}{
		"Habits":  NewHabitDto,
		"Id_user": Id_user,
	}

	h.tmplmain.ExecuteTemplate(w, "main.html", data)
}
func (h *HabitHandler) AddHabitHandler(w http.ResponseWriter, r *http.Request) {
	Id_user, err := h.Auth.Cookie_userID(w, r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	Habit_name := r.FormValue("Habit_name")
	Status_Today := r.FormValue("Status_today") == "on"
	Streak, err := strconv.Atoi(r.FormValue(""))
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	NewHabit := service.HabitResponse{
		Habit_Name:   Habit_name,
		Status_Today: Status_Today,
		Streak:       Streak,
	}

	err = h.HabitService.AddHabitS(NewHabit, Id_user)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func (h *HabitHandler) DeleteHabitHandelr(w http.ResponseWriter, r *http.Request) {
	Id_user, err := h.Auth.Cookie_userID(w, r)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}
	IdHabit, err := strconv.Atoi(r.FormValue("Id"))
	err = h.HabitService.DeleteHabitS(IdHabit, Id_user)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func (h *HabitHandler) UpdateStatusHabit(w http.ResponseWriter, r *http.Request) {
	Id_user, err := h.Auth.Cookie_userID(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	if r.Method == "POST" {
		Id, err := strconv.Atoi(r.FormValue("Id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		streak, err := h.HabitService.GetStreakHabit(Id_user, Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = h.HabitService.ChangeStatusHabit(Id, streak, Id_user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
