package repository

import (
	"github.com/vmihailenco/msgpack/v5"
)

// Prefixes for different entity types in the database
const (
	ServicePrefix = "service:"
	RoutePrefix   = "route:"
)

// makeKey creates a prefixed key for the database
func makeKey(prefix, id string) []byte {
	return []byte(prefix + id)
}

// serialize converts an object to JSON bytes
func serialize(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// deserialize converts JSON bytes to an object
func deserialize(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
