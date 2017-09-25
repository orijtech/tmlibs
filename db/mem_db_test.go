package db_test

import (
	"bytes"
	"testing"

	"github.com/tendermint/tmlibs/db"
)

func TestMemDBClose(t *testing.T) {
	db := db.NewMemDB()
	k, v := []byte("foo"), []byte("bar")
	db.Set(k, v)
	if g, w := db.Get(k), v; !bytes.Equal(g, w) {
		t.Fatalf("got =%x\nwant=%x", g, w)
	}
	db.Close()
	if g := db.Get(k); !bytes.Equal(g, nil) {
		t.Fatalf("After close\ngot =%x\nwant=ni", g)
	}
}
