// internal/models/tariff.go
package models

type Tariff struct {
	ID    int     `db:"id"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}
