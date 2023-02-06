package api

import (
	"api-cv/model"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)


func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}
		sessionToken:= c.Value

		claims := &model.JWTclaim{}
		x, x1 := jwt.ParseWithClaims(sessionToken, claims, func(t *jwt.Token) (interface{}, error) {
			return string(model.JWT_KEY), nil
		})
		_, err = x, x1
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *API) Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Put(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Method is not allowed!"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
