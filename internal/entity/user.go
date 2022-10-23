package entity

// TODO change something db or entity - IsConfirmTerms

type User struct {
	Id             int    `db:"id"`
	UserId         int64  `db:"user_id"`
	Username       string `db:"username"`
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	ChatId         int64  `db:"chat_id"`
	ChatType       string `db:"chat_type"`
	IsConfirmTerms bool   `db:"is_terms_confirmed"`
}
