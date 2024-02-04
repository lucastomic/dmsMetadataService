package idgenerator

import (
	"sync"
	"testing"
)

func TestNewIDGenerator(t *testing.T) {
	generator := New()
	if generator.lastID != 0 {
		t.Errorf("Expected initial lastID to be 0, got %d", generator.lastID)
	}
}

func TestGenerateID(t *testing.T) {
	generator := New()
	firstID := generator.GenerateID()
	if firstID != 1 {
		t.Errorf("Expected first generated ID to be 1, got %d", firstID)
	}

	secondID := generator.GenerateID()
	if secondID != 2 {
		t.Errorf("Expected second generated ID to be 2, got %d", secondID)
	}
}

func TestGenerateID_Concurrency(t *testing.T) {
	generator := New()
	var wg sync.WaitGroup
	ids := make(map[int64]bool)
	idsMutex := sync.Mutex{}
	goroutines := 100
	idsPerGoroutine := 10

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < idsPerGoroutine; j++ {
				id := generator.GenerateID()
				idsMutex.Lock()
				if _, exists := ids[id]; exists {
					t.Errorf("Duplicate ID generated: %d", id)
				}
				ids[id] = true
				idsMutex.Unlock()
			}
		}()
	}
	wg.Wait()

	expectedIDs := goroutines * idsPerGoroutine
	if len(ids) != expectedIDs {
		t.Errorf("Expected %d unique IDs, got %d", expectedIDs, len(ids))
	}
}
