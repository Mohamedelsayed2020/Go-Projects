package controllers

import (
	"Go-Projects/model"
	"Go-Projects/model/common"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	user := model.NewUser()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := common.Conn().Create(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
}

func ListUsers(w http.ResponseWriter, r *http.Request)  {
	var user []model.User
	data := common.Conn().Find(&user)
	if data.Error != nil {
		json.NewEncoder(w).Encode(data.Error)
		return
	}
	json.NewEncoder(w).Encode(data)
}
func GetUser(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id := params["id"]
	user := model.NewUser()
	result := common.Conn().Where("id = ?", id).First(&user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id := params["id"]
	user := model.NewUser()
	result := common.Conn().Where("id = ?", id).Find(&user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	result = common.Conn().Save(user)
	if result.Error != nil {
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	json.NewEncoder(w).Encode(user)
}