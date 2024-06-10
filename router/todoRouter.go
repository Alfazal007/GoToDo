package router

import (
	"net/http"
	"todoapp/controllers"

	"github.com/go-chi/chi"
)

func RouteTodo(apiConf *controllers.ApiConf) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/get-todo", apiConf.GetToDo)
	r.Put("/update-todo", apiConf.UpdateToDo)
	r.Post("/create-todo", controllers.VerifyJWT(apiConf, http.HandlerFunc(apiConf.CreateNewToDo)).ServeHTTP)
	return r
}
