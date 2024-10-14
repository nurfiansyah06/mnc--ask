package tblusers

import (
	"database/sql"
	"fmt"
	"golang-test/model/users"
	"golang-test/repository"
	"time"

	"github.com/google/uuid"
)

type userRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

// Register implements repository.TblUsers.
func (t *userRepositoryImpl) Register(user users.User) (users.User, error) {
	user.UserID = uuid.New().String()
	user.CreatedAt = time.Now()

	query := `INSERT INTO users (user_id, first_name, last_name, phone_number, address, pin, created_at) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := t.DB.Exec(query, user.UserID, user.FirstName, user.LastName, user.PhoneNumber, user.Address, user.Pin, user.CreatedAt)
	if err != nil {
		return users.User{}, err
	}

	return user, nil
}

// Login implements repository.UserRepository.
func (t *userRepositoryImpl) Login(phoneNumber, pin string) (*users.User, error) {
	var user users.User

	query := "SELECT user_id, phone_number, pin FROM users WHERE phone_number = ? AND pin = ?"
	err := t.DB.QueryRow(query, phoneNumber, pin).Scan(&user.UserID, &user.PhoneNumber, &user.Pin)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, fmt.Errorf("no user found")
		}
		return &user, fmt.Errorf("failed to fetch user: %w", err)
	}

	return &user, nil
}

func (t *userRepositoryImpl) SelectByUserId(userId string) (*users.User, error) {
	var user users.User

	query := "SELECT user_id, first_name, last_name, phone_number, address, pin, created_at, balance FROM users WHERE user_id = ?"
	err := t.DB.QueryRow(query, userId).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.Address, &user.Pin, &user.CreatedAt, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, fmt.Errorf("no user found")
		}
		return &user, fmt.Errorf("failed to fetch user: %w", err)
	}

	return &user, nil
}
