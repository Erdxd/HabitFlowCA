package handlers

import (
	"HabitFlow/internal/domain/service"
	"HabitFlow/internal/http/dto"
	"HabitFlow/internal/http/middleware"
	"net/http"
	"text/template"
)

type ProfileHandler struct {
	ProfileService *service.UserService
	Auth           *middleware.JWTToken
	tmplmain       *template.Template
}

func NewProfileHandler(UserService *service.UserService, Auth *middleware.JWTToken, tmpl *template.Template) *ProfileHandler {
	return &ProfileHandler{ProfileService: UserService, Auth: Auth, tmplmain: tmpl}
}
func (P *ProfileHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
	Id_user, err := P.Auth.Cookie_userID(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == "GET" {

		UserData, err := P.ProfileService.GetDataUser(Id_user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(UserData) == 0 {
			http.Error(w, "Пользователя не существует", 401)
			return
		}
		d := UserData[0]
		UserDto := dto.UserRegisterResponse{
			Username: d.Username,
			Email:    d.Email,
		}
		data := map[string]interface{}{
			"Data":    UserDto,
			"Id_user": Id_user,
		}

		P.tmplmain.ExecuteTemplate(w, "profile.html", data)
	}

}
func (P *ProfileHandler) RedactLoginHandler(w http.ResponseWriter, r *http.Request) {
	Id_user, err := P.Auth.Cookie_userID(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {

		newusername := r.FormValue("new_login")
		password := r.FormValue("password")
		err := P.ProfileService.RedactLogin(Id_user, newusername, password)
		if err != nil {
			http.Error(w, "Wrong Password", 401)
			return
		}

		http.Redirect(w, r, "/profile", http.StatusSeeOther)

	}
	P.tmplmain.ExecuteTemplate(w, "redact.html", nil)
}
func (P *ProfileHandler) RedactPassword(w http.ResponseWriter, r *http.Request) {
	Id_user, err := P.Auth.Cookie_userID(w, r)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}
	currentpassword := r.FormValue("Oldpassword")
	NewPassword := r.FormValue("Newpassword")

	err = P.ProfileService.RedactPassword(NewPassword, Id_user, currentpassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
