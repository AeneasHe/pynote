package db

import (
	"testing"
)

func TestDB(t *testing.T) {
	db := NewDataBase("./foo.db", "r+")
	db.Query()
}
