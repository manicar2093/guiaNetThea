package models

type CatalogModel struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type Catalogable interface {
	// GetCatalogInstance crea una instancia de CatalogModel
	GetCatalogInstance() CatalogModel
}
