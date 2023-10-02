package function

import (
	"fmt"
	"net/http"
)

func (s *Server) HandleUsers(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		{
			return s.HandleGetUser(w, r)
		}
	default:
		{
			return WriteJSON(w, http.StatusMethodNotAllowed, fmt.Sprintf("Method %v not allowed", r.Method))
		}
	}
}

func (s *Server) HandleGetUser(w http.ResponseWriter, r *http.Request) error {
	users := []string{"user1", "user2"}
	return WriteJSON(w, http.StatusOK, users)
}
