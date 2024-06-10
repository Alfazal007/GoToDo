package router

import (
	"net/http"
	"todoapp/controllers"

	"github.com/go-chi/chi"
)

func RouteUser(apiConf *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create-user", apiConf.CreateNewUser)
	r.Post("/login", apiConf.Login)
	r.Get("/get-user", controllers.VerifyJWT(apiConf, http.HandlerFunc(apiConf.GetUserFromId)).ServeHTTP)
	return r
}
