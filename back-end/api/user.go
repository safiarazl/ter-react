package api

import (
	"api-cv/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	var creds model.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
		return
	}

	if api.usersRepo.CheckPassLength(creds.Password) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Please provide a password of more than 5 characters"})
		return
	}

	if api.usersRepo.CheckPassAlphabet(creds.Password) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Please use Password with Contains non Alphabetic Characters"})
		return
	}

	err = api.usersRepo.AddUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Username: creds.Username, Message: "User Registered"})
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	var creds model.User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	if creds.Username == "" || creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
		return
	}

	if api.usersRepo.CheckPassLength(creds.Password) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Please provide a password of more than 5 characters"})
		return
	}

	if api.usersRepo.CheckPassAlphabet(creds.Password) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Please use Password with Contains non Alphabetic Characters"})
		return
	}

	err = api.usersRepo.UserAvail(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	expiresAt := time.Now().Add(5 * time.Hour)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	claims := &model.JWTclaim{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "api-cv",
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(model.JWT_KEY)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Path:    "/",
		Value:   token,
		Expires: expiresAt,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Username: creds.Username, Message: "Login Success"})
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	sessionToken := c.Value

	api.sessionsRepo.DeleteSessions(sessionToken)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Username: username, Message: "Logout Success"})
}

func (api *API) SessionValid(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Username: username, Message: "Token Valid"})
}
