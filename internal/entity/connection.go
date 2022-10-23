package entity

import (
	"database/sql"
)

type Connection struct {
	Id              int          `db:"id"`
	Location        string       `db:"location"`
	Port            uint         `db:"port"`
	UserId          int          `db:"user_id"`
	EncryptedSecret string       `db:"encrypted_secret"`
	IpAddress       string       `db:"ip_address"`
	ServerId        int          `db:"server_id"`
	IsActive        bool         `db:"is_active"`
	LastActivate    sql.NullTime `db:"last_activate"`
}

type ConnectionPortCount struct {
	Port  uint `db:"port"`
	Count int  `db:"count"`
}
