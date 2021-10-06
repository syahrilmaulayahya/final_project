package users_test

import (
	"context"
	"final_project/app/middleware"
	"final_project/business/users"
	"final_project/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository mocks.Repository
var userService users.UseCase
var userDomain users.Domain
var reviewDomain users.Review_RatingDomain
var userToken middleware.ConfigJWT

func setup() {
	userService = users.NewUserUseCase(&userRepository, time.Hour*1, userToken)
	userDomain = users.Domain{
		ID:           1,
		Name:         "Syahril",
		Email:        "syahril@gmail.com",
		Password:     "Syahril123",
		Token:        "123",
		Phone_number: 85641441299,
		Gender:       "Male",
		Dob:          time.Now(),
		Address:      "Pekalongan",
		Picture_url:  "www.google.com",
	}
	reviewDomain = users.Review_RatingDomain{
		ID:        1,
		Review:    "Good",
		Rating:    4.5,
		UserID:    1,
		ProductID: 1,
	}
}

func TestRegister(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		userRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.Domain{
			Name:         "Syahril",
			Email:        "syahril@gmail.com",
			Password:     "Syahril123",
			Phone_number: 85641441269,
			Gender:       "Male",
			Dob:          time.Now(),
			Address:      "Pekalongan",
		})
		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Empty Name", func(t *testing.T) {
		userRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.Domain{
			Name:         "",
			Email:        "syahril@gmail.com",
			Password:     "Syahril123",
			Phone_number: 85641441269,
			Gender:       "Male",
			Dob:          time.Now(),
			Address:      "Pekalongan",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Invalid Email", func(t *testing.T) {
		userRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.Domain{
			Name:         "Syahril",
			Email:        "syahrilgmail.com",
			Password:     "Syahril123",
			Phone_number: 85641441269,
			Gender:       "Male",
			Dob:          time.Now(),
			Address:      "Pekalongan",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Invalid Password", func(t *testing.T) {
		userRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.Domain{
			Name:         "Syahril",
			Email:        "syahril@gmail.com",
			Password:     "syahril123",
			Phone_number: 85641441269,
			Gender:       "Male",
			Dob:          time.Now(),
			Address:      "Pekalongan",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 4 | Zero PhoneNumber", func(t *testing.T) {
		userRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.Domain{
			Name:         "Syahril",
			Email:        "syahril@gmail.com",
			Password:     "Syahril123",
			Phone_number: 0,
			Gender:       "Male",
			Dob:          time.Now(),
			Address:      "Pekalongan",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 5 | Invalid Gender", func(t *testing.T) {
		userRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.Domain{
			Name:         "Syahril",
			Email:        "syahril@gmail.com",
			Password:     "Syahril123",
			Phone_number: 85641441269,
			Gender:       "a",
			Dob:          time.Now(),
			Address:      "Pekalongan",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 5 | Empty Dob", func(t *testing.T) {
		userRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.Domain{
			Name:         "Syahril",
			Email:        "syahril@gmail.com",
			Password:     "Syahril123",
			Phone_number: 85641441269,
			Gender:       "Male",
			Address:      "Pekalongan",
		})
		assert.NotNil(t, err)
	})
	t.Run("Test Case 6 | Empty Address", func(t *testing.T) {
		userRepository.On("Register",
			mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.Domain{
			Name:         "Syahril",
			Email:        "syahril@gmail.com",
			Password:     "Syahril123",
			Phone_number: 85641441269,
			Gender:       "Male",
			Dob:          time.Now(),
			Address:      "",
		})
		assert.NotNil(t, err)
	})

}

func TestLogin(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		user, err := userService.Login(context.Background(), "syahril@gmail.com", "syahril123")
		assert.Nil(t, err)
		assert.Equal(t, "Syahril", user.Name)
	})

	t.Run("Test Case 2 | Empty Email", func(t *testing.T) {
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		_, err := userService.Login(context.Background(), "", "syahril123")
		assert.NotNil(t, err)

	})
	t.Run("Test Case 3 |  Empty Password", func(t *testing.T) {
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(userDomain, nil).Once()

		_, err := userService.Login(context.Background(), "syahril@gmail.com", "")
		assert.NotNil(t, err)

	})
}

func TestDetails(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Details", func(t *testing.T) {
		userRepository.On("Details",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		user, err := userService.Details(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, "Syahril", user.Name)
	})
}

func TestUploadReview(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid Review", func(t *testing.T) {
		userRepository.On("UploadReview",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("int")).Return(reviewDomain, nil).Once()

		_, err := userService.UploadReview(context.Background(), users.Review_RatingDomain{
			Review: "Good",
			Rating: 4.5,
		}, 1)
		assert.Nil(t, err)
	})
	setup()
	t.Run("Test Case 2 | Invalid Rating", func(t *testing.T) {
		userRepository.On("UploadReview",
			mock.Anything,
			mock.Anything,
			mock.AnythingOfType("int")).Return(reviewDomain, nil).Once()

		_, err := userService.UploadReview(context.Background(), users.Review_RatingDomain{
			Review: "Good",
			Rating: 7,
		}, 1)
		assert.NotNil(t, err)
	})
}
