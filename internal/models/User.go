package models

import "mime/multipart"

type User struct {
	Id           int    `json:id`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	ProfilePhoto string `json:"profile_photo"`
}

type UserCreateRequest struct {
	Name         string                `form:"name"`
	Lastname     string                `form:"lastname"`
	ProfilePhoto *multipart.FileHeader `form:"profile_photo"`
}

type UserUpdateRequest struct {
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	ProfilePhoto string `json:"profile_photo"`
}
 