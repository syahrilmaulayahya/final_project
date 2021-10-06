package admins_test

import (
	"context"
	"final_project/app/middleware"
	"final_project/business/admins"
	"final_project/business/admins/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var adminRepository mocks.Repository
var adminSevice admins.UseCase
var adminDomain admins.AdminDomain
var adminToken middleware.ConfigJWT

func setup() {
	adminSevice = admins.NewAdminUseCase(&adminRepository, time.Hour*1, adminToken)
	adminDomain = admins.AdminDomain{
		ID:       1,
		Name:     "Syahril",
		Email:    "syahril@gmail.com",
		Password: "syahril123",
		Token:    "123",
	}
}

func TestLogin(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		adminRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(adminDomain, nil).Once()

		admin, err := adminSevice.Login(context.Background(), "syahril@gmail.com", "syahril123")

		assert.Nil(t, err)
		assert.Equal(t, "Syahril", admin.Name)
	})

	t.Run("Test Case 2 | Email Empty", func(t *testing.T) {
		adminRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(adminDomain, nil).Once()

		_, err := adminSevice.Login(context.Background(), "", "syahril123")

		assert.NotNil(t, err)

	})

	t.Run("Test Case 3 | Password Empty", func(t *testing.T) {
		adminRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(adminDomain, nil).Once()

		_, err := adminSevice.Login(context.Background(), "syahril@gmail.com", "")

		assert.NotNil(t, err)

	})

}

func TestRegister(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		adminRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(adminDomain, nil).Once()

		_, err := adminSevice.Register(context.Background(), admins.AdminDomain{
			Name:     "Syahril",
			Email:    "Syahril1@gmail.com",
			Password: "Syahril123",
		})

		assert.Nil(t, err)

	})

	t.Run("Test Case 2 | Name Empty", func(t *testing.T) {
		adminRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(adminDomain, nil).Once()

		_, err := adminSevice.Register(context.Background(), admins.AdminDomain{
			Name:     "",
			Email:    "Syahril1@gmail.com",
			Password: "Syahril123",
		})

		assert.NotNil(t, err)

	})
	t.Run("Test Case 3 | Invalid Email", func(t *testing.T) {
		adminRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(adminDomain, nil).Once()

		_, err := adminSevice.Register(context.Background(), admins.AdminDomain{
			Name:     "Syahril",
			Email:    "Syahril1gmail.com",
			Password: "Syahril123",
		})

		assert.NotNil(t, err)

	})
	t.Run("Test Case 3 | Invalid Email", func(t *testing.T) {
		adminRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(adminDomain, nil).Once()

		_, err := adminSevice.Register(context.Background(), admins.AdminDomain{
			Name:     "Syahril",
			Email:    "Syahril1@gmail.com",
			Password: "syahril123",
		})

		assert.NotNil(t, err)

	})
}
