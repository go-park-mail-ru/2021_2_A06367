package models

//TODO: сгенерировать easyjson файл

type User struct {
	Id int `json:"-"`
	Login string `json:"login"`
	Email string `json:"email"`
	EncryptedPassword string `json:"password"`
}
