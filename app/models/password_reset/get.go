package passwordreset

import (
	"gin_weibo/database"
)

// GetByEmail -
func GetByEmail(email string) (*PasswordReset, error) {
	p := &PasswordReset{}
	d := database.DB.Where("email = ?", email).First(&p)
	return p, d.Error
}

// GetByToken -
func GetByToken(token string) (*PasswordReset, error) {
	p := &PasswordReset{}
	d := database.DB.Where("token = ?", token).First(&p)
	return p, d.Error
}
