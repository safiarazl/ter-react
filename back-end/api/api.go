package api

import (
	repo "api-cv/repository"
	"fmt"
	"net/http"
)

type API struct {
	usersRepo    repo.UserRepository
	sessionsRepo repo.SessionsRepository
	educationsRepo repo.EducationRepository
	mux          *http.ServeMux
}

func NewAPI(usersRepo repo.UserRepository, sessionsRepo repo.SessionsRepository, educationsRepo repo.EducationRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo,
		sessionsRepo,
		educationsRepo,
		mux,
	}

	fileServer := http.FileServer(http.Dir("./template"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))
	// mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))
	// mux.Handle("/user/session/valid", api.Get(api.Auth(http.HandlerFunc(api.SessionValid))))
	// mux.Handle("/user/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))

	// mux.Handle("/user/img/profile", api.Get(api.Auth(http.HandlerFunc(api.ImgProfileView))))
	// mux.Handle("/user/img/update-profile", api.Post(api.Auth(http.HandlerFunc(api.ImgProfileUpdate))))

	// mux.Handle("/product/list", api.Get(api.Auth(http.HandlerFunc(api.ProductList))))
	// mux.Handle("/product/add", api.Post(api.Auth(http.HandlerFunc(api.AddProduct))))
	// mux.Handle("/product/delete", api.Delete(api.Auth(http.HandlerFunc(api.DeleteProduct))))
	// mux.Handle("/product/update", api.Put(api.Auth(http.HandlerFunc(api.UpdateProduct))))

	// mux.Handle("/cart/list", api.Get(api.Auth(http.HandlerFunc(api.CartList))))
	// mux.Handle("/cart/add", api.Post(api.Auth(http.HandlerFunc(api.AddCart))))
	// mux.Handle("/cart/delete", api.Delete(api.Auth(http.HandlerFunc(api.DeleteCart))))
	// mux.Handle("/cart/update", api.Put(api.Auth(http.HandlerFunc(api.UpdateCart))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}