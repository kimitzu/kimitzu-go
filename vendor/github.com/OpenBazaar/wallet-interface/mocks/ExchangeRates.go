// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ExchangeRates is an autogenerated mock type for the ExchangeRates type
type ExchangeRates struct {
	mock.Mock
}

// GetAllRates provides a mock function with given fields: cacheOK
func (_m *ExchangeRates) GetAllRates(cacheOK bool) (map[string]float64, error) {
	ret := _m.Called(cacheOK)

	var r0 map[string]float64
	if rf, ok := ret.Get(0).(func(bool) map[string]float64); ok {
		r0 = rf(cacheOK)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]float64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(cacheOK)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExchangeRate provides a mock function with given fields: currencyCode
func (_m *ExchangeRates) GetExchangeRate(currencyCode string) (float64, error) {
	ret := _m.Called(currencyCode)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(currencyCode)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(currencyCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestRate provides a mock function with given fields: currencyCode
func (_m *ExchangeRates) GetLatestRate(currencyCode string) (float64, error) {
	ret := _m.Called(currencyCode)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(currencyCode)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(currencyCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnitsPerCoin provides a mock function with given fields:
func (_m *ExchangeRates) UnitsPerCoin() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}