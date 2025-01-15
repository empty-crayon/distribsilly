package web

import (
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"net/http"

	"github.com/empty-crayon/distribsilly/db"
)

// Server implements HTTP method handlers used to interact with database
type Server struct {
	db         *db.Database
	shardIdx   int
	shardCount int
	addrs      map[int]string
}

// NewServer creates a new instance http handlers to be used for get and set
func NewServer(db *db.Database, shardIdx, shardCount int, addrs map[int]string) *Server {
	return &Server{
		db:         db,
		shardIdx:   shardIdx,
		shardCount: shardCount,
		addrs:      addrs,
	}
}

func (s *Server) redirect(shard int, w http.ResponseWriter, r *http.Request) {
	targetURL := "http://" + s.addrs[shard] + r.RequestURI
	resp, err := http.Get(targetURL)
	fmt.Fprintf(w, "Redirecting from current shard: %d to target shard: %d (%q)\n", s.shardIdx, shard, targetURL)
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
	shard := s.getShard(key)
	// check if we need to access a different shard
	if shard != s.shardIdx {
		s.redirect(shard, w, r)
		return
	}
	value, err := s.db.GetKey(key)
	if err != nil {
		log.Fatalf("error getin")
	}

	fmt.Fprintf(w, "Shard = %d, Current Shard = %d, addrs of target shard: %q Value: %q, error = %v", shard, s.shardIdx, s.addrs[shard], value, err)
}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")
	value := r.Form.Get("value")
	shard := s.getShard(key)

	if shard != s.shardIdx {
		s.redirect(shard, w, r)
		return
	}

	err := s.db.SetKey(key, []byte(value))

	fmt.Fprintf(w, "Shard: %d, Current Shard: %d, Error = %v", shard, s.shardIdx, err)
}

func (s *Server) getShard(key string) int {
	hash := fnv.New64()
	hash.Write([]byte(key))
	shardId := int(hash.Sum64() % uint64(s.shardCount))
	return shardId
}
