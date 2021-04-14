package fs

import (
	"testing"
)

func TestWalk(t *testing.T) {
	names := Walk("file")
	for _, folder := range names {
		t.Log(folder)
	}
}
