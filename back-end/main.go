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
		DatabaseName: "test_db_camp",
		Port:         5432,
		Schema:       "public",
	}

	conn, err := db.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.User{}, &model.Session{}, &model.Education{})
	usersRepo := repo.NewUserRepository(conn)
	sessionsRepo := repo.NewSessionsRepository(conn)
	educationsRepo := repo.NewEducationRepository(conn)

	mainAPI := api.NewAPI(usersRepo, sessionsRepo, educationsRepo)
	mainAPI.Start()
}
