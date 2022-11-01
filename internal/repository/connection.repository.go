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
	var cp entity.ConnectionPortCount
	query := `SELECT port, count(port) FROM connections
			  GROUP BY port
			  ORDER BY port
			  DESC LIMIT 1;`
	if err := c.db.GetContext(ctx, &cp, query); err != nil {
		c.log.WithContextReqId(ctx).
			Errorf("error get last port from db: %s", err.Error())
		return nil, err
	}
	return &cp, nil
}

func (c *ConnectionRepository) CreateConnection(ctx context.Context, conn *entity.Connection) (int, error) {
	var id int
	query := `INSERT INTO connections (port, encrypted_secret, user_id, server_id)
			  VALUES ($1, $2, $3, $4)
			  RETURNING id;`
	row := c.db.QueryRowxContext(ctx, query, conn.Port, conn.EncryptedSecret, conn.UserId, conn.ServerId)
	if err := row.Scan(&id); err != nil {
		c.log.WithContextReqId(ctx).
			Errorf("error when creating conn: %s, conn: %v", err.Error(), conn)
		return 0, err
	}
	return id, nil
}

func (c *ConnectionRepository) GetConnectionsByUserId(ctx context.Context, id int64) ([]entity.Connection, error) {
	var connections []entity.Connection
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
