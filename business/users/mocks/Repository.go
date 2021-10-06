// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	users "final_project/business/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Details provides a mock function with given fields: ctx, id
func (_m *Repository) Details(ctx context.Context, id int) (users.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) users.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, email, password
func (_m *Repository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	ret := _m.Called(ctx, email, password)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) users.Domain); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, domain
func (_m *Repository) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(context.Context, users.Domain) users.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, users.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadReview provides a mock function with given fields: ctx, domain, id
func (_m *Repository) UploadReview(ctx context.Context, domain users.Review_RatingDomain, id int) (users.Review_RatingDomain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 users.Review_RatingDomain
	if rf, ok := ret.Get(0).(func(context.Context, users.Review_RatingDomain, int) users.Review_RatingDomain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(users.Review_RatingDomain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, users.Review_RatingDomain, int) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
