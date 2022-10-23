package entity

type Server struct {
	Id        int    `db:"id"`
	IpAddress string `db:"ip_address"`
	Location  string `db:"location"`
}
