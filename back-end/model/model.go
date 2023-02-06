package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type Education struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);unique_index"`
	Major string `json:"major"`
	Years string `json:"years"`
}
type Porto struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);unique_index"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Link        string `json:"link"`
}

type Session struct {
	gorm.Model
	Token    string
	Username string `gorm:"type:varchar(100);unique_index"`
	Expiry   time.Time
}

type User struct {
	gorm.Model
	Password string
	Username string `gorm:"type:varchar(100);unique_index"`
}

var JWT_KEY = []byte("qPlAoWkSiJeDuRhFyTg019238747356")

type JWTclaim struct {
	Username string
	jwt.RegisteredClaims
}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
