package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todoapp/helper"
	"todoapp/internal/database"

	"github.com/google/uuid"
)

// register the user
func (apiCfg *ApiConf) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Age      int32  `json:"age"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON", err))
		return
	}
	hashedPassword, err := helper.HashPassword(params.Password)
	if err != nil {
		helper.RespondWithError(w, 400, "There was an issue hashing the password")
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:     params.Name,
		Age:      params.Age,
		ID:       uuid.New(),
		Password: hashedPassword,
	})
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Could not create user %v", err))
		return
	}
	helper.RespondWithJSON(w, 201, helper.CustomUserConvertor(user))
}

// get the user
func (apiCfg *ApiConf) GetUserFromId(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helper.RespondWithError(w, 400, "Issue with finding the user from the database")
		return
	}
	helper.RespondWithJSON(w, 200, helper.CustomUserConvertor(user))
}

func (apiCfg *ApiConf) Login(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("There was an error with the request body sent %v", err))
		return
	}
	user, err := apiCfg.DB.GetUserByName(r.Context(), params.Name)
	if err != nil {
		helper.RespondWithJSON(w, 404, "USER NOT FOUND")
		return
	}
	isValid := helper.CheckPasswordHash(params.Password, user.Password)
	if !isValid {
		helper.RespondWithJSON(w, 400, "Incorrect password")
		return
	}
	// send the api key as well in the headers
	jwtToken, err := GenerateJWT(user)
	if err != nil {
		fmt.Println("The error is ", err)
		helper.RespondWithError(w, 400, "Error generating the token")
		return
	}
	cookie := http.Cookie{
		Name:     "access-token",
		Value:    jwtToken,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	type AccessToken struct {
		AccessToken string `json:"access-token"`
	}
	helper.RespondWithJSON(w, 200, AccessToken{AccessToken: jwtToken})
}
