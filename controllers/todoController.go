package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todoapp/helper"
	"todoapp/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *ApiConf) CreateNewToDo(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helper.RespondWithError(w, 400, "Issue with finding the user from the database")
		return
	}

	type parameters struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON", err))
		return
	}
	todo, err := apiCfg.DB.CreateTodo(r.Context(), database.CreateTodoParams{
		ID:          uuid.New(),
		Title:       params.Title,
		Description: params.Description,
		UserID:      user.ID,
	})
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Could not create user %v", err))
		return
	}
	helper.RespondWithJSON(w, 201, todo)
}

func (apiCfg *ApiConf) GetToDo(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Id uuid.UUID `json:"todo_id"`
	}
	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON", err))
		return
	}
	todo, err := apiCfg.DB.GetTodoById(r.Context(), uuid.UUID(params.Id))
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Could not create user %v", err))
		return
	}
	helper.RespondWithJSON(w, 200, todo)
}

func (apiCfg *ApiConf) UpdateToDo(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		TodoId uuid.UUID `json:"todo_id"`
	}
	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON", err))
		return
	}
	todo, err := apiCfg.DB.UpdateTodo(r.Context(), params.TodoId)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Could not create user %v", err))
		return
	}
	helper.RespondWithJSON(w, 200, todo)
}
