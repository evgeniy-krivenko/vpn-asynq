package repository

import (
	"context"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/entity"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/logger"
	"github.com/jmoiron/sqlx"
)

type ConnectionRepository struct {
	db  *sqlx.DB
	log logger.Logger
}

func (c *ConnectionRepository) GetLastConnectionPortCount(ctx context.Context) (*entity.ConnectionPortCount, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ConnectionRepository) CreateConnection(ctx context.Context, connection *entity.Connection) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ConnectionRepository) GetConnectionsByUserId(ctx context.Context, id int64) ([]entity.Connection, error) {
	var connections []entity.Connection
	// sql syntax
	query := `SELECT connections.id,
					 server_id, port, u.user_id, encrypted_secret,
					 s.ip_address as "ip_address", s.location as "location"
				FROM connections
					LEFT JOIN servers s on s.id = connections.server_id
				    LEFT JOIN users u on connections.user_id = u.id
			    WHERE u.user_id=$1;`
	err := c.db.SelectContext(ctx, &connections, query, id)
	if err != nil {
		return nil, err
	}

	return connections, nil
}

func (c *ConnectionRepository) GetConnectionById(ctx context.Context, id int) (*entity.Connection, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ConnectionRepository) GetConnectionByServerId(ctx context.Context, id int) (*entity.Connection, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ConnectionRepository) SaveConnection(ctx context.Context, conn *entity.Connection) error {
	//TODO implement me
	panic("implement me")
}

func NewConnectionRepository(db *sqlx.DB, log logger.Logger) *ConnectionRepository {
	return &ConnectionRepository{db: db, log: log}
}
