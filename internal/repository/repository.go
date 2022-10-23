package repository

import (
	"context"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/entity"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/logger"
	"github.com/jmoiron/sqlx"
)

type Connection interface {
	GetLastConnectionPortCount(ctx context.Context) (*entity.ConnectionPortCount, error)
	CreateConnection(ctx context.Context, connection *entity.Connection) (int, error)
	GetConnectionsByUserId(ctx context.Context, id int64) ([]entity.Connection, error)
	GetConnectionById(ctx context.Context, id int) (*entity.Connection, error)
	GetConnectionByServerId(ctx context.Context, id int) (*entity.Connection, error)
	SaveConnection(ctx context.Context, conn *entity.Connection) error
}

type Repository struct {
	Connection
}

func NewRepository(db *sqlx.DB, log logger.Logger) *Repository {
	return &Repository{
		Connection: NewConnectionRepository(db, log),
	}
}
