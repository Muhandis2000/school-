// internal/models/payment.go
package models

import "time"

type Payment struct {
	ID        int       `db:"id"`
	StudentID int       `db:"student_id"`
	TariffID  int       `db:"tariff_id"`
	Amount    float64   `db:"amount"`
	Date      time.Time `db:"date"`
}
