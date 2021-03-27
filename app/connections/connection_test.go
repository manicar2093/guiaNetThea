package connections

import "testing"

func TestConnection(t *testing.T) {
	if DB == nil {
		t.Fatal("La base de datos no puede ser nil")
	}
}
