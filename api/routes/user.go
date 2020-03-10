package routes

import (
	"Go-Projects/api/controllers"
	"github.com/gorilla/mux"
	"net/http"
)
func Routes()  {
	r := mux.NewRouter()
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user", controllers.ListUsers).Methods("Get")
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("Get")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")

	http.ListenAndServe(":8000", r)
}