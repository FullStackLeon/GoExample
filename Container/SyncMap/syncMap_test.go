package syncMap

import (
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	var m sync.Map

	// Test Store and Load
	m.Store("key1", "value1")
	if v, ok := m.Load("key1"); !ok || v != "value1" {
		t.Errorf("expected value1, got %v", v)
	}

	// Test LoadOrStore
	if v, loaded := m.LoadOrStore("key2", "value2"); loaded || v != "value2" {
		t.Errorf("expected value2, got %v", v)
	}
	if v, loaded := m.LoadOrStore("key2", "new_value2"); !loaded || v != "value2" {
		t.Errorf("expected value2, got %v", v)
	}

	// Test Delete
	m.Delete("key1")
	if _, ok := m.Load("key1"); ok {
		t.Errorf("expected key1 to be deleted")
	}

	// Test Range
	m.Store("key3", "value3")
	m.Store("key4", "value4")
	keys := make(map[interface{}]bool)
	m.Range(func(key, value interface{}) bool {
		keys[key] = true
		return true
	})
	if len(keys) != 3 {
		t.Errorf("expected 3 keys, got %d", len(keys))
	}
}
