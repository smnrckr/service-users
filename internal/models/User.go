package models

type User struct {
	Id       int    `json:id`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type UserCreateRequest struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type UserUpdateRequest struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}
