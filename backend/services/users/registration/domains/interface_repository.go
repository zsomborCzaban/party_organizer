package domains

import "time"

type IRegistrationRepository interface {
	FindByUsername(string) (*RegistrationRequest, error)
	FindById(id uint) (*RegistrationRequest, error)
	Create(*RegistrationRequest) error
	Delete(reg *RegistrationRequest) error
	DeleteOlderThan(deleteTime time.Time) error
}
