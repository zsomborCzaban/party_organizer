package usecases

import (
	"github.com/stretchr/testify/mock"
	"github.com/zsomborCzaban/party_organizer/services/users/registration/domains"
	"time"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) FindByUsername(username string) (*domains.RegistrationRequest, error) {
	args := mr.Called(username)
	return args.Get(0).(*domains.RegistrationRequest), args.Error(1)
}
func (mr *MockRepository) FindById(id uint) (*domains.RegistrationRequest, error) {
	args := mr.Called(id)
	return args.Get(0).(*domains.RegistrationRequest), args.Error(1)

}
func (mr *MockRepository) Create(regReq *domains.RegistrationRequest) error {
	args := mr.Called(regReq)
	return args.Error(0)
}
func (mr *MockRepository) Delete(regReq *domains.RegistrationRequest) error {
	args := mr.Called(regReq)
	return args.Error(0)
}
func (mr *MockRepository) DeleteOlderThan(deleteTime time.Time) error {
	args := mr.Called(deleteTime)
	return args.Error(0)
}
