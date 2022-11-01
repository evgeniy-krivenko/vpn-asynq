package service

import (
	"context"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/entity"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/repository"
)

const maxClientsOnPortCount = 30

type ConnectionService struct {
	repo   repository.Connection
	crypto Crypto
}

func (c *ConnectionService) CreateConnection(ctx context.Context, usr *entity.User, serverId int) (int, error) {
	connectionPort, err := c.repo.GetLastConnectionPortCount(ctx)
	if err != nil {
		return 0, err
	}

	if connectionPort.Count >= maxClientsOnPortCount {
		connectionPort.Port += 1
	}
	panic("implement me")
}

func (c *ConnectionService) GetConnections(ctx context.Context, userId int64) ([]entity.Connection, error) {
	return c.repo.GetConnectionsByUserId(ctx, userId)
}

func NewConnectionService(repo repository.Connection, crypto Crypto) *ConnectionService {
	return &ConnectionService{repo: repo, crypto: crypto}
}
