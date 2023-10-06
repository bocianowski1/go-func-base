package function

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
	db         DB
}

func NewServer(listenAddr string, db DB) *Server {
	return &Server{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *Server) Start(pairs []HandlerFuncPair) {
	router := mux.NewRouter()
	for _, pair := range pairs {
		router.HandleFunc(pair.Route, pair.Handler)
	}
	http.ListenAndServe(s.listenAddr, router)
}

func (s *Server) Shutdown() {}

func WriteJSON(w http.ResponseWriter, statusCode int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

func MakeHTTPHandler(f HTTPHandler, wp *WorkerPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wp.Enqueue(r)
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, &HTTPError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
		}
	}
}
