package service

import (
	entity "Latihan-Register/Entity"
	"errors"
	"strings"
)

type UserIface interface {
	Register(user *entity.User) (*entity.User, error)
}

type UserSvc struct{}

func NewUserSvc() UserIface {
	return &UserSvc{}
}
func (*UserSvc) Register(user *entity.User) (*entity.User, error) {
	//validasi email
	if !strings.Contains(user.Email, "@gmail.com") {
		return nil, errors.New("Masukan Format Email yg benar")
	}
	if user.Username == "" {
		return nil, errors.New("Tolong isi username")
	}
	if len(user.Password) <= 6 {
		return nil, errors.New("Tolong buat password lebih dari 6")
	}
	if user.Password == "" {
		return nil, errors.New("Tolong buat password lebih dari 6")
	}
	if user.Age <= 8 {
		return nil, errors.New("Tolong usia diatas 8 tahun")
	}

	return user, nil
}
