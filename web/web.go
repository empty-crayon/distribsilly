package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/empty-crayon/distribsilly/db"
)

// Server implements HTTP method handlers used to interact with database
type Server struct {
	db         *db.Database
	shardIdx   int
	shardCount int
}

// NewServer creates a new instance http handlers to be used for get and set
func NewServer(db *db.Database, shardIdx, shardCount int) *Server {
	return &Server{
		db:         db,
		shardIdx:   shardIdx,
		shardCount: shardCount,
	}
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")

	value, err := s.db.GetKey(key)
	if err != nil {
		log.Fatalf("error getin")
	}
	fmt.Fprintf(w, "Value: %q, error = %v", value, err)
}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")
	value := r.Form.Get("value")

	err := s.db.SetKey(key, []byte(value))

	fmt.Fprintf(w, "Error = %v", err)
}
