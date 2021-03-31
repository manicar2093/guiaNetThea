package dao

import (
	"testing"

	"github.com/manicar2093/guianetThea/app/connections"
)

func TestFindEndpointByName(t *testing.T) {
	dao := NewEndpointDao(connections.DB)
	name := "inicio"
	endpoint, e := dao.FindEndpointByName(name)

	if e != nil {
		t.Fatal("No debió haber error: ", e)
	}

	if endpoint.EndpointID <= 0 {
		t.Fatal("No hay ID en la estructura")
	}
}

func TestFindEndpointByID(t *testing.T) {
	dao := NewEndpointDao(connections.DB)
	id := int32(1)
	endpoint, e := dao.FindEndpointByID(id)

	if e != nil {
		t.Fatal("No debió haber error: ", e)
	}

	if endpoint.Name == "" {
		t.Fatal("No hay Name en la estructura")
	}
}
