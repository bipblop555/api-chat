package models

import (
	"net/http"
	"time"

	"fmt"

	"App/internal/resources"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

type User struct {
	Id        int        `gorm:"column:id;primaryKey"`
	Email     string     `gorm:"not null;unique_index"`
	Username  string     `gorm:"not null;unique_index"`
	Password  string     `gorm:"no null"` // Ne pas store dans la database
	CreatedAt *time.Time `gorm:"type:timestamp"`
	UpdatedAt *time.Time `gorm:"type:timestamp;autoUpdateTime:true"`
	UserLogs  []UserLog  `gorm:"foreignKey:UserID"`
}

type FindUser struct {
	Username string `gorm:"not null;unique_index"`
}

type UserLog struct {
	Id           int `gorm:"primaryKey"`
	LogTimestamp *time.Time
	LogData      datatypes.JSON `gorm:"type:jsonb"`
	UserID       uint           `gorm:"column:user_id;not null"`
}

// ByEmail pour get user by email
func (ug *DbGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.Db.Table("users").Where("email = ?", email).First(&user)

	err := first(db, &user)
	return &user, err
}

func (ug *DbGorm) ByUserName(username string) ([]resources.UserResource, error) {
	var user []resources.UserResource
	db := ug.Db.Table("users").Where("username LIKE ?", username+"%").Limit(50).Find(&user)

	fmt.Print(user)

	//err := first(db, &user)
	if db.Error != nil {
		return nil, db.Error
	}

	return user, nil
}

// Methode Create pour add user to db
func (ug *DbGorm) Create(entity interface{}, w http.ResponseWriter) error {
	db := ug.Db.Table("users").Create(entity)

	if db.Error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return db.Error
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}

// Authenticate Method is used for Authenticate and Validate login
func (us *DatabaseProvider) Authenticate(email, password string) (*User, error) {
	// Vlidate if the user is existed in the database or no
	foundUser, err := us.ByEmail(email)

	fmt.Println("foundUser => ", foundUser)

	if err != nil {
		return nil, err
	}

	// Compare the login based in the Hash value
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))

	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		case nil:
			return nil, err
		default:
			return nil, err
		}
	}

	return foundUser, nil
}

func (ug *DbGorm) GetAllUsers() ([]User, error) {
	var users []User

	db := ug.Db.Table("users").Order("firstname").Find(&users)
	if db.Error != nil {
		return nil, db.Error
	}

	return users, nil
}
