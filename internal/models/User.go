package models

type User struct {
	Id           int    `json:id`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	ProfilePhoto string `json:"profile_photo"`
}

type UserCreateRequest struct {
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	ProfilePhoto string `json:"profile_photo"`
}

type UserUpdateRequest struct {
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	ProfilePhoto string `json:"profile_photo"`
}
