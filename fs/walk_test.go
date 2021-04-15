package fs

import (
	"testing"
)

func TestWalk(t *testing.T) {
	path := "/Users/aeneas/Github/Cofepy/youdao"
	names := ShowPath(path, "file")
	for _, folder := range names {
		t.Log(folder)
	}
}
