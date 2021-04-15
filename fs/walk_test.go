package fs

import (
	"testing"
)

func TestWalk(t *testing.T) {
	path := "/Users/aeneas/Github/Cofepy/youdao"
	names := ShowPath(path, "all")
	for _, folder := range names {
		t.Log(folder)
	}
}
