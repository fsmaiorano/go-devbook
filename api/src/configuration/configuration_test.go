package configuration

import (
	"testing"
)

// CreateUser creates a new user
func TestLoad(t *testing.T) {
	t.Run("Configuration load", func(t *testing.T) {
		t.Log("Starting UnitTest Configuration load")

		Load()

		if ApiPort == 0 {
			t.Error("ApiPort is not set")
		}

		if ConnectionString == "sqlserver://%s:%s@%s:%s?database=%s" {
			t.Error("ConnectionString is not set")
		}
	})
}
