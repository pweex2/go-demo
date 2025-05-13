package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) FindByID(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func TestFindUser(t *testing.T) {
	mockRepo := new(MockUserRepo)
	expected := &User{ID: 1, Name: "Wayne"}

	mockRepo.On("FindByID", 1).Return(expected, nil)

	user, err := mockRepo.FindByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "Wayne", user.Name)
	mockRepo.AssertExpectations(t)
}
