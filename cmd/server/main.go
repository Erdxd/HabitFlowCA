package main

import (
	"HabitFlow/internal/domain/service"
	"HabitFlow/internal/http/handlers"
	"HabitFlow/internal/http/middleware"
	"HabitFlow/internal/infrastructure/database"
	"HabitFlow/internal/infrastructure/hashing"
	infrastructure "HabitFlow/internal/infrastructure/jwt"
	"HabitFlow/internal/infrastructure/scheduler"
	"HabitFlow/internal/infrastructure/telegram"
	"HabitFlow/internal/repository"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	tmpl, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Println(err)
	}

	urlDb := os.Getenv("DATABASE_URL")

	db, err := database.InitDb(urlDb)
	if err != nil {
		log.Println(err)
	}
	hasher := hashing.NewHashService()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewuserService(userRepo, hasher)
	tokentg := os.Getenv("TELEGRAM_BOT_TOKEN")

	jwt := os.Getenv("JWT_TOKEN")

	jwtservice := infrastructure.NewJWTService(jwt)
	jwttoken := middleware.NewJwtKey(jwtservice)
	jwtService2 := service.NewTokenService(jwtservice)
	AdminRepository := repository.NewAdminRepo(db)
	AdminService := service.NewAdminService(AdminRepository, hasher)
	AdminHandler := handlers.NewAdminHandler(AdminService, jwttoken, tmpl, jwtService2)
	authMiddleWare := middleware.NewAuthMiddleware(jwtservice)

	habitRepository := repository.NewHabitRepository(db)
	habitService := service.NewHabitService(habitRepository)
	habitHanlder := handlers.NewHabitHandler(habitService, jwttoken, tmpl)
	UserHanlder := handlers.NewUserHandler(userService, jwttoken, tmpl, jwtService2)
	Profilehandler := handlers.NewProfileHandler(userService, jwttoken, tmpl)
	cron := scheduler.NewScheduler()
	cron.ResetStatus("00:00", func() { habitService.ResetAllStatusHabit() })
	cron.Start()
	telegrambot := telegram.NewBot(tokentg, userService, habitService)
	go telegrambot.HandleMessages()

	http.HandleFunc("/", habitHanlder.CheckHabit)
	http.HandleFunc("/register", UserHanlder.Register)
	http.HandleFunc("/login", UserHanlder.Login)
	http.HandleFunc("/profile", Profilehandler.ProfileHandler)
	http.HandleFunc("/redactlogin", Profilehandler.RedactLoginHandler)
	http.HandleFunc("/redactpassword", Profilehandler.RedactPassword)
	http.HandleFunc("/admin/users", authMiddleWare.AdminOnly(AdminHandler.CheckUsers))
	http.HandleFunc("/admin/update-user", authMiddleWare.AdminOnly(AdminHandler.UpdatePassword))
	http.HandleFunc("/admin/delete-user", authMiddleWare.AdminOnly(AdminHandler.DeleteUser))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
