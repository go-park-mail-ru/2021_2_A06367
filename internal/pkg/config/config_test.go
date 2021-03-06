package config

import (
	"os"
	"testing"
)

func TestCheckConfig(t *testing.T) {
	os.Setenv("DATABASE_URL", "value")
	_, err := GetConnectionString()
	if err != nil {
		t.Error("not env")
	}
}

func TestCheckConfigError(t *testing.T) {
	os.Setenv("DATABASE_URL", "")
	_, err := GetConnectionString()
	if err == nil {
		t.Error("not env")
	}
}
