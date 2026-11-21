// Package api is the HTTP-free entry point over the store.
package api

import "github.com/coupling-dev/archfit-coldrun/internal/store"

// Server wraps a store for callers.
type Server struct {
	s *store.Store
}

// New returns a Server with an empty store.
func New() *Server {
	return &Server{s: store.New()}
}

// Put stores a value under key.
func (srv *Server) Put(key, value string) error {
	return srv.s.Set(key, value)
}

// Get returns the value under key, or "" when absent.
func (srv *Server) Get(key string) (string, bool) {
	return srv.s.Lookup(key)
}
