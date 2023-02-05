package repository

import (
	"api-cv/model"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	return u.db.Create(&session).Error //TODO: replace this
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	return u.db.Where("token = ?", tokenTarget).Delete(&model.Session{}).Error //TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	return u.db.Model(&session).Where("username = ?", session.Username).Updates(&session).Error //TODO: replace this
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	var session model.Session
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return session, err
	}
	if u.TokenExpired(session) {
		return model.Session{}, gorm.ErrRecordNotFound
	}
	return session, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	var session model.Session
	err := u.db.Where("username = ?", name).First(&session).Error // TODO: replace this

	return session, err // TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := u.db.Where("token = ?", token).First(&session).Error
	return session, err // TODO: replace this
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
