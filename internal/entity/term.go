package entity

import "time"

type Term struct {
	Id          int        `db:"id"`
	OrderNumber int        `db:"order_number"`
	Text        string     `db:"text"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
