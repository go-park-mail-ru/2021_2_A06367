package config

import (
	"fmt"
	"testing"
)

func TestFilmsRepo_GetFilmById(t *testing.T) {
	conn, err := GetConnectionString()
	if err != nil {
		t.Error("not env")
	}
	fmt.Println(conn)
}
