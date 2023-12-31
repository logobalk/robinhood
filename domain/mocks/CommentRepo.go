// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	domain "robinhood/domain"

	mock "github.com/stretchr/testify/mock"
)

// CommentRepo is an autogenerated mock type for the CommentRepo type
type CommentRepo struct {
	mock.Mock
}

// GetAllCommentByAppId provides a mock function with given fields: appId
func (_m *CommentRepo) GetAllCommentByAppId(appId string) ([]*domain.Comment, error) {
	ret := _m.Called(appId)

	var r0 []*domain.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*domain.Comment, error)); ok {
		return rf(appId)
	}
	if rf, ok := ret.Get(0).(func(string) []*domain.Comment); ok {
		r0 = rf(appId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(appId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveComment provides a mock function with given fields: comment
func (_m *CommentRepo) SaveComment(comment *domain.Comment) error {
	ret := _m.Called(comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Comment) error); ok {
		r0 = rf(comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCommentRepo creates a new instance of CommentRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommentRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommentRepo {
	mock := &CommentRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
