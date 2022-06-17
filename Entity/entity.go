package entity

import "time"

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Age       int
	Create_at time.Time
	Update_at time.Time
}
