package model

import (
	"gorm.io/gorm"
	"time"
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
	Name     string  `gorm:"type:varchar(100);unique_index"`
	Major     string  `json:"major"`
	Years     string  `json:"years"`
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

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}