// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	user "authapp/domain/user"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetUserByPhonenumber provides a mock function with given fields: phonenumber
func (_m *Repository) GetUserByPhonenumber(phonenumber string) (*user.User, error) {
	ret := _m.Called(phonenumber)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(phonenumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phonenumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUserPass provides a mock function with given fields: phonenumber, password
func (_m *Repository) GetUserByUserPass(phonenumber string, password string) (*user.User, error) {
	ret := _m.Called(phonenumber, password)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string, string) *user.User); ok {
		r0 = rf(phonenumber, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(phonenumber, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Persist provides a mock function with given fields: u
func (_m *Repository) Persist(u *user.User) (*user.User, error) {
	ret := _m.Called(u)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(*user.User) *user.User); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
