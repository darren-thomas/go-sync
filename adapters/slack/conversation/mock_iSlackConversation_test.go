// Code generated by mockery. DO NOT EDIT.

package conversation

import (
	slack "github.com/slack-go/slack"
	mock "github.com/stretchr/testify/mock"
)

// mockISlackConversation is an autogenerated mock type for the iSlackConversation type
type mockISlackConversation struct {
	mock.Mock
}

type mockISlackConversation_Expecter struct {
	mock *mock.Mock
}

func (_m *mockISlackConversation) EXPECT() *mockISlackConversation_Expecter {
	return &mockISlackConversation_Expecter{mock: &_m.Mock}
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *mockISlackConversation) GetUserByEmail(email string) (*slack.User, error) {
	ret := _m.Called(email)

	var r0 *slack.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*slack.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *slack.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slack.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockISlackConversation_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type mockISlackConversation_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - email string
func (_e *mockISlackConversation_Expecter) GetUserByEmail(email interface{}) *mockISlackConversation_GetUserByEmail_Call {
	return &mockISlackConversation_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", email)}
}

func (_c *mockISlackConversation_GetUserByEmail_Call) Run(run func(email string)) *mockISlackConversation_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *mockISlackConversation_GetUserByEmail_Call) Return(_a0 *slack.User, _a1 error) *mockISlackConversation_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockISlackConversation_GetUserByEmail_Call) RunAndReturn(run func(string) (*slack.User, error)) *mockISlackConversation_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsersInConversation provides a mock function with given fields: params
func (_m *mockISlackConversation) GetUsersInConversation(params *slack.GetUsersInConversationParameters) ([]string, string, error) {
	ret := _m.Called(params)

	var r0 []string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(*slack.GetUsersInConversationParameters) ([]string, string, error)); ok {
		return rf(params)
	}
	if rf, ok := ret.Get(0).(func(*slack.GetUsersInConversationParameters) []string); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(*slack.GetUsersInConversationParameters) string); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(*slack.GetUsersInConversationParameters) error); ok {
		r2 = rf(params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// mockISlackConversation_GetUsersInConversation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsersInConversation'
type mockISlackConversation_GetUsersInConversation_Call struct {
	*mock.Call
}

// GetUsersInConversation is a helper method to define mock.On call
//   - params *slack.GetUsersInConversationParameters
func (_e *mockISlackConversation_Expecter) GetUsersInConversation(params interface{}) *mockISlackConversation_GetUsersInConversation_Call {
	return &mockISlackConversation_GetUsersInConversation_Call{Call: _e.mock.On("GetUsersInConversation", params)}
}

func (_c *mockISlackConversation_GetUsersInConversation_Call) Run(run func(params *slack.GetUsersInConversationParameters)) *mockISlackConversation_GetUsersInConversation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*slack.GetUsersInConversationParameters))
	})
	return _c
}

func (_c *mockISlackConversation_GetUsersInConversation_Call) Return(_a0 []string, _a1 string, _a2 error) *mockISlackConversation_GetUsersInConversation_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *mockISlackConversation_GetUsersInConversation_Call) RunAndReturn(run func(*slack.GetUsersInConversationParameters) ([]string, string, error)) *mockISlackConversation_GetUsersInConversation_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsersInfo provides a mock function with given fields: users
func (_m *mockISlackConversation) GetUsersInfo(users ...string) (*[]slack.User, error) {
	_va := make([]interface{}, len(users))
	for _i := range users {
		_va[_i] = users[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *[]slack.User
	var r1 error
	if rf, ok := ret.Get(0).(func(...string) (*[]slack.User, error)); ok {
		return rf(users...)
	}
	if rf, ok := ret.Get(0).(func(...string) *[]slack.User); ok {
		r0 = rf(users...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]slack.User)
		}
	}

	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(users...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockISlackConversation_GetUsersInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsersInfo'
type mockISlackConversation_GetUsersInfo_Call struct {
	*mock.Call
}

// GetUsersInfo is a helper method to define mock.On call
//   - users ...string
func (_e *mockISlackConversation_Expecter) GetUsersInfo(users ...interface{}) *mockISlackConversation_GetUsersInfo_Call {
	return &mockISlackConversation_GetUsersInfo_Call{Call: _e.mock.On("GetUsersInfo",
		append([]interface{}{}, users...)...)}
}

func (_c *mockISlackConversation_GetUsersInfo_Call) Run(run func(users ...string)) *mockISlackConversation_GetUsersInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *mockISlackConversation_GetUsersInfo_Call) Return(_a0 *[]slack.User, _a1 error) *mockISlackConversation_GetUsersInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockISlackConversation_GetUsersInfo_Call) RunAndReturn(run func(...string) (*[]slack.User, error)) *mockISlackConversation_GetUsersInfo_Call {
	_c.Call.Return(run)
	return _c
}

// InviteUsersToConversation provides a mock function with given fields: channelID, users
func (_m *mockISlackConversation) InviteUsersToConversation(channelID string, users ...string) (*slack.Channel, error) {
	_va := make([]interface{}, len(users))
	for _i := range users {
		_va[_i] = users[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, channelID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *slack.Channel
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...string) (*slack.Channel, error)); ok {
		return rf(channelID, users...)
	}
	if rf, ok := ret.Get(0).(func(string, ...string) *slack.Channel); ok {
		r0 = rf(channelID, users...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slack.Channel)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(channelID, users...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockISlackConversation_InviteUsersToConversation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InviteUsersToConversation'
type mockISlackConversation_InviteUsersToConversation_Call struct {
	*mock.Call
}

// InviteUsersToConversation is a helper method to define mock.On call
//   - channelID string
//   - users ...string
func (_e *mockISlackConversation_Expecter) InviteUsersToConversation(channelID interface{}, users ...interface{}) *mockISlackConversation_InviteUsersToConversation_Call {
	return &mockISlackConversation_InviteUsersToConversation_Call{Call: _e.mock.On("InviteUsersToConversation",
		append([]interface{}{channelID}, users...)...)}
}

func (_c *mockISlackConversation_InviteUsersToConversation_Call) Run(run func(channelID string, users ...string)) *mockISlackConversation_InviteUsersToConversation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *mockISlackConversation_InviteUsersToConversation_Call) Return(_a0 *slack.Channel, _a1 error) *mockISlackConversation_InviteUsersToConversation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockISlackConversation_InviteUsersToConversation_Call) RunAndReturn(run func(string, ...string) (*slack.Channel, error)) *mockISlackConversation_InviteUsersToConversation_Call {
	_c.Call.Return(run)
	return _c
}

// KickUserFromConversation provides a mock function with given fields: channelID, user
func (_m *mockISlackConversation) KickUserFromConversation(channelID string, user string) error {
	ret := _m.Called(channelID, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(channelID, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockISlackConversation_KickUserFromConversation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'KickUserFromConversation'
type mockISlackConversation_KickUserFromConversation_Call struct {
	*mock.Call
}

// KickUserFromConversation is a helper method to define mock.On call
//   - channelID string
//   - user string
func (_e *mockISlackConversation_Expecter) KickUserFromConversation(channelID interface{}, user interface{}) *mockISlackConversation_KickUserFromConversation_Call {
	return &mockISlackConversation_KickUserFromConversation_Call{Call: _e.mock.On("KickUserFromConversation", channelID, user)}
}

func (_c *mockISlackConversation_KickUserFromConversation_Call) Run(run func(channelID string, user string)) *mockISlackConversation_KickUserFromConversation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *mockISlackConversation_KickUserFromConversation_Call) Return(_a0 error) *mockISlackConversation_KickUserFromConversation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockISlackConversation_KickUserFromConversation_Call) RunAndReturn(run func(string, string) error) *mockISlackConversation_KickUserFromConversation_Call {
	_c.Call.Return(run)
	return _c
}

// newMockISlackConversation creates a new instance of mockISlackConversation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockISlackConversation(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockISlackConversation {
	mock := &mockISlackConversation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
