// Code generated by mockery. DO NOT EDIT.

package team

import (
	context "context"

	github "github.com/google/go-github/v47/github"
	mock "github.com/stretchr/testify/mock"
)

// mockIGitHubTeam is an autogenerated mock type for the iGitHubTeam type
type mockIGitHubTeam struct {
	mock.Mock
}

type mockIGitHubTeam_Expecter struct {
	mock *mock.Mock
}

func (_m *mockIGitHubTeam) EXPECT() *mockIGitHubTeam_Expecter {
	return &mockIGitHubTeam_Expecter{mock: &_m.Mock}
}

// AddTeamMembershipBySlug provides a mock function with given fields: ctx, org, slug, user, opts
func (_m *mockIGitHubTeam) AddTeamMembershipBySlug(ctx context.Context, org string, slug string, user string, opts *github.TeamAddTeamMembershipOptions) (*github.Membership, *github.Response, error) {
	ret := _m.Called(ctx, org, slug, user, opts)

	var r0 *github.Membership
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *github.TeamAddTeamMembershipOptions) (*github.Membership, *github.Response, error)); ok {
		return rf(ctx, org, slug, user, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *github.TeamAddTeamMembershipOptions) *github.Membership); ok {
		r0 = rf(ctx, org, slug, user, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Membership)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *github.TeamAddTeamMembershipOptions) *github.Response); ok {
		r1 = rf(ctx, org, slug, user, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, string, *github.TeamAddTeamMembershipOptions) error); ok {
		r2 = rf(ctx, org, slug, user, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// mockIGitHubTeam_AddTeamMembershipBySlug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTeamMembershipBySlug'
type mockIGitHubTeam_AddTeamMembershipBySlug_Call struct {
	*mock.Call
}

// AddTeamMembershipBySlug is a helper method to define mock.On call
//   - ctx context.Context
//   - org string
//   - slug string
//   - user string
//   - opts *github.TeamAddTeamMembershipOptions
func (_e *mockIGitHubTeam_Expecter) AddTeamMembershipBySlug(ctx interface{}, org interface{}, slug interface{}, user interface{}, opts interface{}) *mockIGitHubTeam_AddTeamMembershipBySlug_Call {
	return &mockIGitHubTeam_AddTeamMembershipBySlug_Call{Call: _e.mock.On("AddTeamMembershipBySlug", ctx, org, slug, user, opts)}
}

func (_c *mockIGitHubTeam_AddTeamMembershipBySlug_Call) Run(run func(ctx context.Context, org string, slug string, user string, opts *github.TeamAddTeamMembershipOptions)) *mockIGitHubTeam_AddTeamMembershipBySlug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(*github.TeamAddTeamMembershipOptions))
	})
	return _c
}

func (_c *mockIGitHubTeam_AddTeamMembershipBySlug_Call) Return(_a0 *github.Membership, _a1 *github.Response, _a2 error) *mockIGitHubTeam_AddTeamMembershipBySlug_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *mockIGitHubTeam_AddTeamMembershipBySlug_Call) RunAndReturn(run func(context.Context, string, string, string, *github.TeamAddTeamMembershipOptions) (*github.Membership, *github.Response, error)) *mockIGitHubTeam_AddTeamMembershipBySlug_Call {
	_c.Call.Return(run)
	return _c
}

// ListTeamMembersBySlug provides a mock function with given fields: ctx, org, slug, opts
func (_m *mockIGitHubTeam) ListTeamMembersBySlug(ctx context.Context, org string, slug string, opts *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error) {
	ret := _m.Called(ctx, org, slug, opts)

	var r0 []*github.User
	var r1 *github.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error)); ok {
		return rf(ctx, org, slug, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *github.TeamListTeamMembersOptions) []*github.User); ok {
		r0 = rf(ctx, org, slug, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*github.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *github.TeamListTeamMembersOptions) *github.Response); ok {
		r1 = rf(ctx, org, slug, opts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string, *github.TeamListTeamMembersOptions) error); ok {
		r2 = rf(ctx, org, slug, opts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// mockIGitHubTeam_ListTeamMembersBySlug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTeamMembersBySlug'
type mockIGitHubTeam_ListTeamMembersBySlug_Call struct {
	*mock.Call
}

// ListTeamMembersBySlug is a helper method to define mock.On call
//   - ctx context.Context
//   - org string
//   - slug string
//   - opts *github.TeamListTeamMembersOptions
func (_e *mockIGitHubTeam_Expecter) ListTeamMembersBySlug(ctx interface{}, org interface{}, slug interface{}, opts interface{}) *mockIGitHubTeam_ListTeamMembersBySlug_Call {
	return &mockIGitHubTeam_ListTeamMembersBySlug_Call{Call: _e.mock.On("ListTeamMembersBySlug", ctx, org, slug, opts)}
}

func (_c *mockIGitHubTeam_ListTeamMembersBySlug_Call) Run(run func(ctx context.Context, org string, slug string, opts *github.TeamListTeamMembersOptions)) *mockIGitHubTeam_ListTeamMembersBySlug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*github.TeamListTeamMembersOptions))
	})
	return _c
}

func (_c *mockIGitHubTeam_ListTeamMembersBySlug_Call) Return(_a0 []*github.User, _a1 *github.Response, _a2 error) *mockIGitHubTeam_ListTeamMembersBySlug_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *mockIGitHubTeam_ListTeamMembersBySlug_Call) RunAndReturn(run func(context.Context, string, string, *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error)) *mockIGitHubTeam_ListTeamMembersBySlug_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveTeamMembershipBySlug provides a mock function with given fields: ctx, org, slug, user
func (_m *mockIGitHubTeam) RemoveTeamMembershipBySlug(ctx context.Context, org string, slug string, user string) (*github.Response, error) {
	ret := _m.Called(ctx, org, slug, user)

	var r0 *github.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*github.Response, error)); ok {
		return rf(ctx, org, slug, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *github.Response); ok {
		r0 = rf(ctx, org, slug, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, org, slug, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockIGitHubTeam_RemoveTeamMembershipBySlug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveTeamMembershipBySlug'
type mockIGitHubTeam_RemoveTeamMembershipBySlug_Call struct {
	*mock.Call
}

// RemoveTeamMembershipBySlug is a helper method to define mock.On call
//   - ctx context.Context
//   - org string
//   - slug string
//   - user string
func (_e *mockIGitHubTeam_Expecter) RemoveTeamMembershipBySlug(ctx interface{}, org interface{}, slug interface{}, user interface{}) *mockIGitHubTeam_RemoveTeamMembershipBySlug_Call {
	return &mockIGitHubTeam_RemoveTeamMembershipBySlug_Call{Call: _e.mock.On("RemoveTeamMembershipBySlug", ctx, org, slug, user)}
}

func (_c *mockIGitHubTeam_RemoveTeamMembershipBySlug_Call) Run(run func(ctx context.Context, org string, slug string, user string)) *mockIGitHubTeam_RemoveTeamMembershipBySlug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *mockIGitHubTeam_RemoveTeamMembershipBySlug_Call) Return(_a0 *github.Response, _a1 error) *mockIGitHubTeam_RemoveTeamMembershipBySlug_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockIGitHubTeam_RemoveTeamMembershipBySlug_Call) RunAndReturn(run func(context.Context, string, string, string) (*github.Response, error)) *mockIGitHubTeam_RemoveTeamMembershipBySlug_Call {
	_c.Call.Return(run)
	return _c
}

// newMockIGitHubTeam creates a new instance of mockIGitHubTeam. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockIGitHubTeam(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockIGitHubTeam {
	mock := &mockIGitHubTeam{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
