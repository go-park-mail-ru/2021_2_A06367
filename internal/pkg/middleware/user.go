package middleware

import "github.com/go-park-mail-ru/2021_2_A06367/internal/models"

//TODO сделать более умную валидацию
func UserIsValid(user models.User) bool {
	return user.Login != "" && user.EncryptedPassword != "" && user.Email != ""
}

func LoginUserIsValid(user models.LoginUser) bool {
	return user.Login != "" && user.EncryptedPassword != ""
}
