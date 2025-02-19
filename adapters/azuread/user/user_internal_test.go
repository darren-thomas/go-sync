package user

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	gosync "github.com/ovotech/go-sync"
)

type MockRequestAdapter struct {
	SerializationWriterFactory serialization.SerializationWriterFactory
}

func (r *MockRequestAdapter) Send(
	_ context.Context,
	_ *abstractions.RequestInformation,
	_ serialization.ParsableFactory,
	_ abstractions.ErrorMappings,
) (serialization.Parsable, error) {
	return nil, nil
}

func (r *MockRequestAdapter) SendEnum(
	_ context.Context,
	_ *abstractions.RequestInformation,
	_ serialization.EnumFactory,
	_ abstractions.ErrorMappings,
) (any, error) {
	return nil, nil //nolint:nilnil
}

func (r *MockRequestAdapter) SendCollection(
	_ context.Context,
	_ *abstractions.RequestInformation,
	_ serialization.ParsableFactory,
	_ abstractions.ErrorMappings,
) ([]serialization.Parsable, error) {
	return nil, nil
}

func (r *MockRequestAdapter) SendEnumCollection(
	_ context.Context,
	_ *abstractions.RequestInformation,
	_ serialization.EnumFactory,
	_ abstractions.ErrorMappings,
) ([]any, error) {
	return nil, nil
}

func (r *MockRequestAdapter) SendPrimitive(
	_ context.Context,
	_ *abstractions.RequestInformation,
	_ string,
	_ abstractions.ErrorMappings,
) (any, error) {
	return nil, nil //nolint:nilnil
}

func (r *MockRequestAdapter) SendPrimitiveCollection(
	_ context.Context,
	_ *abstractions.RequestInformation,
	_ string,
	_ abstractions.ErrorMappings,
) ([]any, error) {
	return nil, nil
}

func (r *MockRequestAdapter) SendNoContent(
	_ context.Context,
	_ *abstractions.RequestInformation,
	_ abstractions.ErrorMappings,
) error {
	return nil
}

func (r *MockRequestAdapter) ConvertToNativeRequest(
	_ context.Context,
	_ *abstractions.RequestInformation,
) (any, error) {
	return nil, nil //nolint:nilnil
}

func (r *MockRequestAdapter) GetSerializationWriterFactory() serialization.SerializationWriterFactory {
	return r.SerializationWriterFactory
}

func (r *MockRequestAdapter) EnableBackingStore(_ store.BackingStoreFactory) {
}

//nolint:revive,stylecheck
func (r *MockRequestAdapter) SetBaseUrl(_ string) {
}

//nolint:revive,stylecheck
func (r *MockRequestAdapter) GetBaseUrl() string {
	return ""
}

func TestTeam_Get(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	mockClient := newMockIClient(t)
	mockUser := newMockIUser(t)

	mockClient.On("Users").Return(&users.UsersRequestBuilder{})
	mockClient.On("GetAdapter").Return(&MockRequestAdapter{})

	adapter := New(mockClient)
	adapter.users = mockUser

	expected := []string{"test.email.1@ovo.com", "test.email.2@ovo.com", "test.email.3@ovo.com"}
	respOut := make([]models.Userable, 0, 3)

	for _, m := range expected {
		t := models.NewUser()
		t.SetMail(to.Ptr(m))
		respOut = append(respOut, t)
	}

	resp := models.NewUserCollectionResponse()
	resp.SetValue(respOut)

	mockUser.On("Get", ctx, mock.Anything).Return(resp, nil)

	emails, err := adapter.Get(ctx)
	if err != nil {
		t.Errorf("Unexpected error. Wanted nil, got: %s", err)

		return
	}

	assert.ElementsMatch(t, expected, emails)
}

func TestTeam_Add(t *testing.T) {
	t.Parallel()

	client := newMockIClient(t)
	client.On("Users").Return(&users.UsersRequestBuilder{})

	adapter := New(client)
	err := adapter.Add(context.TODO(), []string{"foo@email", "bar@email"})

	assert.ErrorIs(t, err, gosync.ErrReadOnly)
}

func TestTeam_Remove(t *testing.T) {
	t.Parallel()

	client := newMockIClient(t)
	client.On("Users").Return(&users.UsersRequestBuilder{})

	adapter := New(client)
	err := adapter.Remove(context.TODO(), []string{"foo@email", "bar@email"})

	assert.ErrorIs(t, err, gosync.ErrReadOnly)
}

func Test_isAdvancedQuery(t *testing.T) {
	t.Parallel()

	type args struct {
		filter string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal query", args{"mail eq 'test.user@ovo.com'"}, false},
		{"ovo email", args{"endswith(mail, '@ovo.com'"}, true},
		{"not specific email", args{"mail ne 'test.user@ovo.com"}, true},
		{"exclude test email", args{"NOT startswith(mail, 'test.user@')"}, true},
		{"company", args{"companyName eq 'OVO'"}, true},
		{"cost center", args{"employeeOrgData/costCenter eq 'tech'"}, true},
		{"division", args{"employeeOrgData/division eq 'infrared'"}, true},
		{"employeeType", args{"employeeType eq 'FTE'"}, true},
		{"officeLocation", args{"officeLocation eq 'SBK'"}, true},
		{"extensionAttribute1", args{"onPremisesExtensionAttributes/extensionAttribute1 eq 'Full'"}, true},
		{"extensionAttribute2", args{"onPremisesExtensionAttributes/extensionAttribute2 eq 'Full'"}, true},
		{"extensionAttribute10", args{"onPremisesExtensionAttributes/extensionAttribute10 eq 'Full'"}, true},
		{"extensionAttribute15", args{"onPremisesExtensionAttributes/extensionAttribute15 eq 'Full'"}, true},
		{"extensionAttribute16", args{"onPremisesExtensionAttributes/extensionAttribute16 eq 'Full'"}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := isAdvancedQuery(tt.args.filter); got != tt.want {
				t.Errorf("isAdvancedQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInit(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		adapter, err := Init(ctx, map[gosync.ConfigKey]string{
			Filter: "filter",
		})

		assert.NoError(t, err)
		assert.IsType(t, &User{}, adapter)
		assert.Equal(t, "filter", adapter.(*User).filter)
	})

	t.Run("missing config", func(t *testing.T) {
		t.Parallel()

		t.Run("missing filter", func(t *testing.T) {
			t.Parallel()

			adapter, err := Init(ctx, map[gosync.ConfigKey]string{})

			assert.NoError(t, err)
			assert.Equal(t, "", adapter.(*User).filter)
		})
	})
}

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		client := newMockIClient(t)
		client.On("Users").Return(&users.UsersRequestBuilder{})

		adapter := New(client)
		assert.Equal(t, client, adapter.client)
	})

	t.Run("with filter config", func(t *testing.T) {
		t.Parallel()

		filter := "mail eq 'test@example.com'"

		client := newMockIClient(t)
		client.On("Users").Return(&users.UsersRequestBuilder{})

		adapter := New(client, WithFilter(filter))
		assert.Equal(t, filter, adapter.filter)
	})
}
