package domains

import "golang.org/x/crypto/bcrypt"

type LoginRequest struct {
	Username *string `json:"username" validate:"required,min=3"`
	Password *string `json:"password" validate:"required"`
}

func (l *LoginRequest) CheckPassword(userPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPwd), []byte(*l.Password))
}
