package models

import (
	"App/internal/modules/hash"
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"net/http"
	// "gorm.io/datatypes"
)

var (
	// ErrNotFound retourné si ressource absente database
	ErrNotFound = errors.New("Models: Resource Not Found")
	// ErrInvalidID utilisée quand on passe un ID à la méthode Delete pour delete un user de la DB
	ErrInvalidID = errors.New("Models: ID must be Valid ID")
	// UserPwPepper ajouté pepper value
	UserPwPepper = "secret-random-string"
	// ErrInvalidPassword pour retourne invalide password
	ErrInvalidPassword = errors.New("Models: Invalid Password")
	// HmacSecret for creating the HMAC
	HmacSecret          = "secret-hmac-key"
	_          EntityDB = &DbGorm{}
)

// UserDB interface handle toute les opérations User dans la DB
// Couche database pour les queries single user

type EntityDB interface {
	// Alter
	Create(entity interface{}, w http.ResponseWriter) error
	Update(entity interface{}, id string, w http.ResponseWriter) error
	Delete(entity interface{}, id string) (DeleteMessage, error)
	ByID(id string, entity interface{}) error
	ByEmail(email string) (*User, error)
	ByUserName(username string) ([]User, error)
	CreateMessage(entity interface{}, w http.ResponseWriter) error
	GetAllLinkedChat(senderID int) ([]MessageChat, error)
	GetAllMessagesFromUser(senderId string, receiverId string) ([]Message, error)
	Close() error
	Ping() error

	GetAllUsers() ([]User, error)
}

// Database Auth Layer
type EntityImplementService interface {
	Authenticate(email, password string) (*User, error)
	EntityDB
}

type EntityService struct {
	EntityDB
	db *gorm.DB
}

type dbConnectionValidator struct {
	EntityDB
	hmac hash.HMAC
}

type DbGorm struct {
	Db    *gorm.DB
	Dbase *sql.DB
}

func first(db *gorm.DB, entity interface{}) error {
	err := db.First(entity).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}

	return err
}

// ByID method to get a user by ID
func (ug *DbGorm) ByID(id string, entity interface{}) error {
	db := ug.Db.Where("id = ?", id).First(entity)
	err := first(db, entity)
	return err
}

// Update method to update a user in database
func (ug *DbGorm) Update(entity interface{}, id string, w http.ResponseWriter) error {

	w.WriteHeader(http.StatusCreated)
	return nil
}
