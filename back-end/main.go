package main

import (
	"api-cv/model"
	"api-cv/db"
	"api-cv/api"
	repo "api-cv/repository"
)

func main () {
	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "4321",
		DatabaseName: "postgres",
		Port:         5432,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.User{}, &model.Session{}, &model.Education{}, &model.Porto{})
	usersRepo := repo.NewUserRepository(conn)
	sessionsRepo := repo.NewSessionsRepository(conn)
	educationsRepo := repo.NewEducationRepository(conn)
	portoRepo := repo.NewPortoRepository(conn)

	mainAPI := api.NewAPI(usersRepo, sessionsRepo, educationsRepo, portoRepo)
	mainAPI.Start()
}
