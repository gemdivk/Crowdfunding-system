package tests

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) CreateUser(email string, hashedPassword string) error {
	args := m.Called(email, hashedPassword)
	return args.Error(0)
}

func (m *MockDB) GetUserByEmail(email string) (string, error) {
	args := m.Called(email)
	return args.String(0), args.Error(1)
}

// User model
type User struct {
	DB *MockDB
}

func (u *User) Register(email, password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	return u.DB.CreateUser(email, hashedPassword)
}

func (u *User) Authenticate(email, password string) (bool, error) {
	storedPassword, err := u.DB.GetUserByEmail(email)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func TestRegister(t *testing.T) {
	mockDB := new(MockDB)
	user := &User{DB: mockDB}
	email := "test@example.com"
	password := "securepassword"

	mockDB.On("CreateUser", email, mock.Anything).Return(nil)

	err := user.Register(email, password)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestAuthenticate_Success(t *testing.T) {
	mockDB := new(MockDB)
	user := &User{DB: mockDB}
	email := "test@example.com"
	password := "securepassword"

	hashedPassword, _ := hashPassword(password)
	mockDB.On("GetUserByEmail", email).Return(hashedPassword, nil)

	ok, err := user.Authenticate(email, password)
	assert.NoError(t, err)
	assert.True(t, ok)
	mockDB.AssertExpectations(t)
}

func TestAuthenticate_Failure(t *testing.T) {
	mockDB := new(MockDB)
	user := &User{DB: mockDB}
	email := "test@example.com"
	password := "wrongpassword"

	hashedPassword, _ := hashPassword("securepassword")
	mockDB.On("GetUserByEmail", email).Return(hashedPassword, nil)

	ok, err := user.Authenticate(email, password)
	assert.NoError(t, err)
	assert.False(t, ok)
	mockDB.AssertExpectations(t)
}

func TestAuthenticate_UserNotFound(t *testing.T) {
	mockDB := new(MockDB)
	user := &User{DB: mockDB}
	email := "nonexistent@example.com"
	password := "securepassword"

	mockDB.On("GetUserByEmail", email).Return("", sql.ErrNoRows)

	ok, err := user.Authenticate(email, password)
	assert.Error(t, err)
	assert.False(t, ok)
	mockDB.AssertExpectations(t)
}

func TestHashPassword(t *testing.T) {
	hashedPassword, err := hashPassword("securepassword")
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}
