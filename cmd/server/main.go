package main

import (
	"HabitFlow/internal/domain/service"
	"HabitFlow/internal/http/handlers"
	"HabitFlow/internal/http/middleware"
	"HabitFlow/internal/infrastructure/database"
	"HabitFlow/internal/infrastructure/hashing"
	infrastructure "HabitFlow/internal/infrastructure/jwt"
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
	tmpl, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		log.Println(err)
	}

	urlDb := os.Getenv("DATABASE_URL")
	log.Println("Database URL:", urlDb)
	db, err := database.InitDb(urlDb)
	if err != nil {
		log.Println(err)
	}
	hasher := hashing.NewHashService()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewuserService(userRepo, hasher)
	jwt := os.Getenv("JWT_TOKEN")
	log.Println("JWT Token:", jwt)
	jwtservice := infrastructure.NewJWTService(jwt)
	jwttoken := middleware.NewJwtKey(jwtservice)
	jwtService2 := service.NewTokenService(jwtservice)
	//authMiddleWare := middleware.NewAuthMiddleware(jwtservice)

	habitRepository := repository.NewHabitRepository(db)
	habitService := service.NewHabitService(habitRepository)
	habitHanlder := handlers.NewHabitHandler(habitService, jwttoken, tmpl)
	UserHanlder := handlers.NewUserHandler(userService, jwttoken, tmpl, jwtService2)

	http.HandleFunc("/", habitHanlder.CheckHabit)
	http.HandleFunc("/register", UserHanlder.Register)
	http.HandleFunc("/login", UserHanlder.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
