package domains

type ChangePasswordRequest struct {
	Password        string `json:"password" validate:"required,min=3,containsany=0123456789"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
