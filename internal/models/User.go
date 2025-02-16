package models

import (
	"errors"
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation"
)

var ErrorNoRowsAffected = errors.New("no rows affected")

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	ProfilePhoto string `json:"profile_photo"`
}

func (User) TableName() string {
	return "users"
}

type UserCreateRequest struct {
	Name         string                `form:"name" json:"name" required:"true"`
	Lastname     string                `form:"lastname" json:"lastname" required:"true"`
	ProfilePhoto *multipart.FileHeader `form:"profile_photo" json:"profile_photo"`
}

func (f UserCreateRequest) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name, validation.Required.Error("name bulunmalı")),
		validation.Field(&f.Lastname, validation.Required.Error("lastname bulunmalı")),
		validation.Field(&f.ProfilePhoto, validation.NilOrNotEmpty.Error("geçersiz profile fotoğrafı")),
	)
}

func (f UserUpdateRequest) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Name, validation.Length(2, 50).Error("name 2-50 karakter arasında olmalı")),
		validation.Field(&f.Lastname, validation.Length(2, 50).Error("lastname 2-50 karakter arasında olmalı")),
		validation.Field(&f.ProfilePhoto, validation.NilOrNotEmpty.Error("geçersiz profile fotoğrafı")),
	)
}

type UserUpdateRequest struct {
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	ProfilePhoto string `json:"profile_photo"`
}
