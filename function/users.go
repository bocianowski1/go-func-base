package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *Server) HandleUsers(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Query().Get("email") != "" {
		email := r.URL.Query().Get("email")
		switch r.Method {
		case "GET":
			{
				return s.HandleGetUserBy(w, r, email)
			}

		case "DELETE":
			{
				return s.HandleDeleteUserBy(w, r, email)
			}
		}
	}
	switch r.Method {
	case "GET":
		{
			return s.HandleGetUser(w, r)
		}
	case "POST":
		{
			return s.HandleCreateUser(w, r)
		}
	default:
		{
			return WriteJSON(w, http.StatusMethodNotAllowed, fmt.Sprintf("Method %v not allowed", r.Method))
		}
	}
}

func (s *Server) HandleGetUser(w http.ResponseWriter, r *http.Request) error {
	users, err := s.db.GetUser()
	if err != nil {
		return err
	}
	if users == nil {
		return WriteJSON(w, http.StatusNotFound, fmt.Sprintf("No users found"))
	}
	return WriteJSON(w, http.StatusOK, users)
}

func (s *Server) HandleGetUserBy(w http.ResponseWriter, r *http.Request, email string) error {
	user, err := s.db.GetUserBy(email)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, HTTPError{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("User with email %v not found", email),
		})
	}

	return WriteJSON(w, http.StatusOK, user)
}

func (s *Server) HandleDeleteUserBy(w http.ResponseWriter, r *http.Request, email string) error {
	err := s.db.DeleteUserBy(email)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, fmt.Sprintf("User with email %v deleted", email))
}

func (s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	req := CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	user := NewUser(req.FirstName, req.LastName, req.Email, req.Password)

	id, err := s.db.CreateUser(user)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	user.ID = id

	return WriteJSON(w, http.StatusOK, user)
}
