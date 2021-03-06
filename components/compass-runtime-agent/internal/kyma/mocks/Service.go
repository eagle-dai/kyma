// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	apperrors "github.com/kyma-project/kyma/components/compass-runtime-agent/internal/apperrors"
	kyma "github.com/kyma-project/kyma/components/compass-runtime-agent/internal/kyma"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-project/kyma/components/compass-runtime-agent/internal/kyma/model"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Apply provides a mock function with given fields: applications
func (_m *Service) Apply(applications []model.Application) ([]kyma.Result, apperrors.AppError) {
	ret := _m.Called(applications)

	var r0 []kyma.Result
	if rf, ok := ret.Get(0).(func([]model.Application) []kyma.Result); ok {
		r0 = rf(applications)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kyma.Result)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func([]model.Application) apperrors.AppError); ok {
		r1 = rf(applications)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
