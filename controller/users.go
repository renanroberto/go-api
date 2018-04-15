package controller

import (
	"api/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var currId int = len(model.Users)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	users := model.Users

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		panic(err)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		params = mux.Vars(r)
		users  = model.Users
		found  = false
	)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		if user.Id == id {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(user)
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("User not found")
	}
}

// type Response struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// 	Email     string `json:"email"`
// 	Password  string `json:"password"`
// }

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	var response model.User

	json.NewDecoder(r.Body).Decode(&response)

	currId++
	user := model.User{
		currId,
		response.Firstname,
		response.Lastname,
		response.Email,
		response.Password,
		true,
	}

	model.Users = append(model.Users, user)

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	var response model.User
	json.NewDecoder(r.Body).Decode(&response)

	user := model.User{
		id,
		response.Firstname,
		response.Lastname,
		response.Email,
		response.Password,
		response.Active,
	}

	for i := range model.Users {
		if model.Users[i].Id == id {
			model.Users[i] = user

			json.NewEncoder(w).Encode(user)

			break
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	for i := range model.Users {
		if model.Users[i].Id == id {
			model.Users[i].Active = false

			json.NewEncoder(w).Encode(model.Users[i])

			break
		}
	}
}
