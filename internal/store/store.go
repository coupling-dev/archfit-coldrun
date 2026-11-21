// Package store keeps values in memory.
package store

import (
	"errors"

	"github.com/coupling-dev/archfit-coldrun/internal/util"
)

// Store is a concurrency-unsafe key/value map.
type Store struct {
	data map[string]string
}

// New returns an empty Store.
func New() *Store {
	return &Store{data: make(map[string]string)}
}

// Set stores value under key after validating the key.
func (s *Store) Set(key, value string) error {
	if !util.ValidKey(key) {
		return errors.New("store: invalid key")
	}
	s.data[key] = value
	return nil
}

// Lookup returns the value under key.
func (s *Store) Lookup(key string) (string, bool) {
	v, ok := s.data[key]
	return v, ok
}
