package main

import (
	entity "Latihan-Register/Entity"
	"time"

	// service "Latihan-Register/Service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var PORT = ":8080"

var mapUser = map[int]entity.User{
	0: {Id: 0, Username: "arip", Email: "arip@gmail.com", Password: "Password123", Age: 23},
	1: {Id: 1, Username: "dasril", Email: "dasril@gmail.com", Password: "Password123", Age: 23},
	2: {Id: 2, Username: "mega", Email: "mega@gmail.com", Password: "Password123", Age: 23},
}

func main() {
	http.HandleFunc("/users/", getEmployeesbyId)
	http.ListenAndServe(PORT, nil)
}
func getEmployeesbyId(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path[1:])
	paths := strings.Split(r.URL.Path[1:], "/")
	fmt.Printf("%+v", paths)
	fmt.Println(len(paths))
	// fmt.Println(users11[1])

	switch r.Method {
	case "GET":
		if paths[1] != "" {
			if idInt, err := strconv.Atoi(paths[1]); err == nil {
				fmt.Println("validasi benar2")
				jsonData, _ := json.Marshal(mapUser[idInt])
				// fmt.Println(&users11[idInt])
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
			}
		} else {
			jsonData, _ := json.Marshal(mapUser)
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
		JsonData, _ := json.Marshal(mapUser)
		w.Header().Add("Content-Type", "application/json")
		w.Write(JsonData)
	case "PUT":
		fmt.Println("PUT")
		var newUser entity.User
		json.NewDecoder(r.Body).Decode(&newUser)
		newUser.Create_at = time.Now()
		newUser.Update_at = time.Now()
		mapUser[int(newUser.Id)] = newUser
		JsonData, _ := json.Marshal(mapUser)
		w.Header().Add("Content-Type", "application/json")
		w.Write(JsonData)
	case "DELETE":
		if paths[1] != "" {
			if idInt, err := strconv.Atoi(paths[1]); err == nil {
				delete(mapUser,idInt)
				jsonData, _ := json.Marshal(mapUser)
				// fmt.Println(&users11[idInt])
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
			}
		}
		// var newUser entity.User
		// json.NewDecoder(r.Body).Decode(&newUser)
		// delete(mapUser, newUser.Id)
		// JsonData, _ := json.Marshal(mapUser)
		// w.Header().Add("Content-Type", "application/json")
		// w.Write(JsonData)

	}

}
