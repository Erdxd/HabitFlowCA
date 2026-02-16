package handlers

import (
	"HabitFlow/internal/domain/service"
	"HabitFlow/internal/http/middleware"
	"log"
	"net/http"
	"text/template"
)

type Userhandler struct {
	UserService *service.UserService
	auth        *middleware.JWTToken
	tmplmain    *template.Template
	jwtservice  *service.TokenService
}

func NewUserHandler(UserHandler *service.UserService, Auth *middleware.JWTToken, tmpl *template.Template, jwt *service.TokenService) *Userhandler {
	return &Userhandler{UserService: UserHandler, auth: Auth, tmplmain: tmpl, jwtservice: jwt}
}
func (us *Userhandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Login := r.FormValue("username")
		Email := r.FormValue("email")
		Password := r.FormValue("password")
		PasswordRep := r.FormValue("password_repeat")

		user := service.UserRegisterResponse{
			Username: Login,
			Email:    Email,
			Password: Password,
		}
		err := us.UserService.Register(user, PasswordRep)
		if err != nil {
			http.Error(w, "Cant register you", 400)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	us.tmplmain.ExecuteTemplate(w, "register.html", nil)
}
func (us *Userhandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Login := r.FormValue("login")
		Password := r.FormValue("password")
		err := us.UserService.Login(Login, Password)

		if err != nil {
			http.Error(w, "wrong login or password", 401)
			return
		}

		ID_user, err := us.UserService.GetUserId(Login)
		if err != nil {
			log.Println(err)
			http.Error(w, "something wrong3", 500)
			return
		}

		Admin, err := us.UserService.GetAdmin(ID_user)
		if err != nil {

			http.Error(w, "something wrong2", 500)
			return
		}
		token, err := us.jwtservice.CreateToken(ID_user, Admin)

		if err != nil {

			http.Error(w, "something wrong1", 500)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			MaxAge:   86400,
			HttpOnly: true,
			Secure:   false,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
	us.tmplmain.ExecuteTemplate(w, "login.html", nil)

}
