package main

import (
	entity "Latihan-Register/Entity"
	"log"
	"time"

	// service "Latihan-Register/Service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var PORT = ":8080"

var mapUser = map[int]entity.User{
	0: {Id: 0, Username: "arip", Email: "arip@gmail.com", Password: "Password123", Age: 23},
	1: {Id: 1, Username: "dasril", Email: "dasril@gmail.com", Password: "Password123", Age: 23},
	2: {Id: 2, Username: "mega", Email: "mega@gmail.com", Password: "Password123", Age: 23},
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/users", getEmployeesbyId)
	r.HandleFunc("/users/{id}", getEmployeesbyId)
	fmt.Println("Now loading on port 0.0.0.0" + PORT)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
func getEmployeesbyId(w http.ResponseWriter, r *http.Request) {
	paths := mux.Vars(r)
	id := paths["id"]
	fmt.Println("")
	switch r.Method {
	case "GET":
		if id != "" {
			if idInt, err := strconv.Atoi(id); err == nil {
				if user, ok := mapUser[idInt]; ok {
					jsonData, _ := json.Marshal(user)
					w.Header().Add("Content-Type", "application/json")
					w.Write(jsonData)
					return
				} else {
					w.Write([]byte("Data not found"))
					return
				}
			}
		} else {
			var sliceUser []entity.User
			for _, v := range mapUser {
				sliceUser = append(sliceUser, v)
			}
			jsonData, _ := json.Marshal(sliceUser)
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
		}
	case "POST":
		fmt.Println("Post")
		var newUser entity.User
		json.NewDecoder(r.Body).Decode(&newUser)
		newUser.Create_at = time.Now()
		newUser.Update_at = time.Now()
		mapUser[int(newUser.Id)] = newUser
		var sliceUser []entity.User
		for _, v := range mapUser {
			sliceUser = append(sliceUser, v)
		}
		jsonData, _ := json.Marshal(sliceUser)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	case "PUT":
		fmt.Println("PUT")
		if id != "" {
			if idInt, err := strconv.Atoi(id); err == nil {
				if _, ok := mapUser[idInt]; ok {
					var newUser entity.User
					json.NewDecoder(r.Body).Decode(&newUser)
					newUser.Create_at = time.Now()
					newUser.Update_at = time.Now()
					mapUser[int(newUser.Id)] = newUser
					var sliceUser []entity.User
					for _, v := range mapUser {
						sliceUser = append(sliceUser, v)
					}
					jsonData, _ := json.Marshal(sliceUser)
					w.Header().Add("Content-Type", "application/json")
					w.Write(jsonData)
					return
				} else {
					w.Write([]byte("Data not found"))
					return
				}
			}
		}
		
	case "DELETE":
		if id != "" {
			if idInt, err := strconv.Atoi(id); err == nil {
				if _, ok := mapUser[idInt]; ok {
					delete(mapUser, idInt)
					var sliceUser []entity.User
					for _, v := range mapUser {
						sliceUser = append(sliceUser, v)
					}
					jsonData, _ := json.Marshal(sliceUser)
					w.Header().Add("Content-Type", "application/json")
					w.Write(jsonData)
					return
				} else {
					w.Write([]byte("Data not found"))
					return
				}
			}
		}
	}

}
