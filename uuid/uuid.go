// Package UUID provides an easy to replace UUID package.
// Forks of Event Horizon can re-implement this package with a UUID library of choice.
package uuid

import "github.com/segmentio/ksuid"

// UUID is an alias type for github.com/google/uuid.UUID.
type UUID = ksuid.KSUID

// Nil is an empty UUID.
var Nil = UUID(ksuid.Nil)

// New creates a new UUID.
func New() UUID {
	return UUID(ksuid.New())
}

// Parse parses a UUID from a string, or returns an error.
func Parse(s string) (UUID, error) {
	id, err := ksuid.Parse(s)

	return UUID(id), err
}

// MustParse parses a UUID from a string, or panics.
func MustParse(s string) UUID {
	id, err := ksuid.Parse(s)
	if err != nil {
		panic(err)
	}
	return UUID(id)
}
