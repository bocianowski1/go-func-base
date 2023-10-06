package function

import (
	"crypto"
	"net/http"
	"time"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

type HTTPError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type HandlerFuncPair struct {
	Route   string
	Handler http.HandlerFunc
}

type User struct {
	ID         string    `bson:"_id,omitempty"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	Password   []byte    `json:"-"`
	Token      string    `json:"token"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

func NewUser(firstName, lastName, email, password string) *User {
	return &User{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Password:   crypto.SHA256.New().Sum([]byte(password)),
		Role:       "USER",
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
	}
}

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
