package server

import (
	"testing"
)

func TestServer(t *testing.T) {
	config := "/Users/aeneas/Github/Cofepy/pynote/pynote/config.json"
	s := NewServer(config)
	s.Run()
}
