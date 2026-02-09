package handlers

import (
	"HabitFlow/internal/domain/models"
	"HabitFlow/internal/domain/service"
	"HabitFlow/internal/http/middleware"
	"text/template"

	"log"
	"net/http"
)

type HabitHandler struct {
	HabitService *service.HabitService
	Auth         *middleware.JWTToken
	tmplmain     *template.Template
}
type Userhandler struct {
	UserService *service.UserService
	auth        *middleware.JWTToken
	tmplmain    *template.Template
}

func NewHabitHandler(HabitService *service.HabitService, Auth *middleware.JWTToken, tmpl *template.Template) *HabitHandler {
	return &HabitHandler{HabitService: HabitService, Auth: Auth, tmplmain: tmpl}
}
func NewUserHandler(HabitHandler *service.UserService, Auth *middleware.JWTToken, tmpl *template.Template) *Userhandler {
	return &Userhandler{UserService: HabitHandler, auth: Auth, tmplmain: tmpl}
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
	h.tmplmain.ExecuteTemplate(w, "index.html", data)
}
