package web

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/empty-crayon/distribsilly/config"
	"github.com/empty-crayon/distribsilly/db"
)

// Server implements HTTP method handlers used to interact with database
type Server struct {
	db     *db.Database
	shards *config.Shards
}

// NewServer creates a new instance http handlers to be used for get and set
func NewServer(db *db.Database, s *config.Shards) *Server {
	return &Server{
		db:     db,
		shards: s,
	}
}

func (s *Server) redirect(shard int, w http.ResponseWriter, r *http.Request) {
	targetURL := "http://" + s.shards.Addrs[shard] + r.RequestURI
	fmt.Fprintf(w, "Redirecting from current shard: %d to target shard: %d (%q)\n", s.shards.CurIdx, shard, targetURL)

	resp, err := http.Get(targetURL)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error redirecting the request: %v", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)

}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")

	shard := s.shards.GetShard(key)
	// check if we need to access a different shard
	if shard != s.shards.CurIdx {
		s.redirect(shard, w, r)
		return
	}

	value, err := s.db.GetKey(key)
	if err != nil {
		log.Fatalf("error getin")
	}

	fmt.Fprintf(w, "Shard = %d, Current Shard = %d, addrs of target shard: %q Value: %q, error = %v", shard, s.shards.CurIdx, s.shards.Addrs[shard], value, err)
}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")
	value := r.Form.Get("value")
	shard := s.shards.GetShard(key)

	if shard != s.shards.CurIdx {
		s.redirect(shard, w, r)
		return
	}

	err := s.db.SetKey(key, []byte(value))

	fmt.Fprintf(w, "Shard: %d, Current Shard: %d, Error = %v", shard, s.shards.CurIdx, err)
}
