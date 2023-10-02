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
	ID         int       `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	Password   []byte    `json:"password"`
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

type Product struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	UserID      int       `json:"userId"`
	CreatedAt   time.Time `json:"createdAt"`
	ModifiedAt  time.Time `json:"modifiedAt"`
}

func NewProduct(title, description string, price, userID int) *Product {
	return &Product{
		Title:       title,
		Description: description,
		Price:       price,
		UserID:      userID,
		CreatedAt:   time.Now().UTC(),
		ModifiedAt:  time.Now().UTC(),
	}
}

type CreateProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	UserID      int    `json:"userId"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
