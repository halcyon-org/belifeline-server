package e2e

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")

	code := m.Run()
	os.Exit(code)
}
