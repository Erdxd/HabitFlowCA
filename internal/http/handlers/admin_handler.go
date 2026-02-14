package handlers

import (
	"HabitFlow/internal/domain/service"
	"HabitFlow/internal/http/dto"
	"HabitFlow/internal/http/middleware"
	"net/http"
	"strconv"
	"text/template"
)

type Adminhandler struct {
	AdminService *service.AdminService
	auth         *middleware.JWTToken
	tmplmain     *template.Template
	jwtservice   *service.TokenService
}

func NewAdminHandler(AdminHandler *service.AdminService, Auth *middleware.JWTToken, tmpl *template.Template, jwt *service.TokenService) *Adminhandler {
	return &Adminhandler{AdminService: AdminHandler, auth: Auth, tmplmain: tmpl, jwtservice: jwt}
}

func (A *Adminhandler) CheckUsers(w http.ResponseWriter, r *http.Request) {
	UsersAll, err := A.AdminService.GetDataAboutAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var usersalldto []dto.AdminDataUsers
	for _, users := range UsersAll {
		Users := dto.AdminDataUsers{
			Id_user:  users.Id_user,
			Username: users.Username,
			Email:    users.Email,
			Password: users.Password,
		}
		usersalldto = append(usersalldto, Users)

	}
	DataAboutUsers := map[string]interface{}{
		"DataAboutUsers": usersalldto,
	}
	A.tmplmain.ExecuteTemplate(w, "admin.html", DataAboutUsers)

}
func (A *Adminhandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	Id_user, err := strconv.Atoi(r.FormValue("Id_user"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	NewPassword := r.FormValue("password")
	err = A.AdminService.ChangePasswordForUser(Id_user, NewPassword)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
func (A *Adminhandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	Id_user, err := strconv.Atoi(r.FormValue("Id_user"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = A.AdminService.DeleteAllHabits(Id_user)
	if err != nil {

		http.Error(w, "Could not delete this account", http.StatusInternalServerError)
		return
	}
	err = A.AdminService.DeleteAccount(Id_user)
	if err != nil {
		http.Error(w, "Could not delete this account", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
