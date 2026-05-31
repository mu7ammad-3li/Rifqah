package room

import (
	"testing"
	"strings"
)

func TestGenerateShortID(t *testing.T) {
	id, err := generateShortID()
	if err != nil {
		t.Fatalf("Failed to generate short ID: %v", err)
	}

	if len(id) != 9 {
		t.Errorf("Expected length 9, got %d (%s)", len(id), id)
	}

	parts := strings.Split(id, "-")
	if len(parts) != 2 {
		t.Errorf("Expected format XXXX-YYYY, got %s", id)
	}

	if len(parts[0]) != 4 || len(parts[1]) != 4 {
		t.Errorf("Expected 4 chars on each side of hyphen, got %d and %d", len(parts[0]), len(parts[1]))
	}
}
