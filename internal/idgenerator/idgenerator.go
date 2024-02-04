package idgenerator

import (
	"sync"
)

// IDGenerator defines an interface for generating unique identifiers.
// The interface is generic, allowing for the generation of identifiers of any type.
// C represents the type of the identifier to be generated.
type IDGenerator[C any] interface {
	// GenerateID generates and returns a new unique identifier of type C.
	GenerateID() C
}

// idGenerator provides a thread-safe implementation of the IDGenerator interface for generating
// unique int64 identifiers. It is designed to be used in environments where unique sequential
// identifiers are needed, such as database entries, file storage, or anywhere a unique reference
// is required.
type idGenerator struct {
	mutex  sync.Mutex // mutex protects access to lastID, ensuring thread-safe operations.
	lastID int64      // lastID holds the value of the last generated ID.
}

// New initializes and returns a new instance of idGenerator.
// The returned idGenerator starts generating IDs from 1, ensuring that the first
// generated ID is always 1.
func New() idGenerator {
	return idGenerator{
		mutex:  sync.Mutex{},
		lastID: 0, // Start IDs from 1 by initializing lastID to 0.
	}
}

// GenerateID increments and returns the next unique identifier.
// It ensures thread safety by locking around the increment operation,
// making it safe to use across multiple goroutines.
func (g *idGenerator) GenerateID() int64 {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.lastID++
	return g.lastID
}
