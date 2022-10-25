package service

import (
	"context"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/entity"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/repository"
)

type ConnectionService struct {
	repo   repository.Connection
	crypto Crypto
}

func (c ConnectionService) CreateConnection(ctx context.Context, usr *entity.User, serverId int) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c ConnectionService) GetConnections(ctx context.Context, userId int64) ([]entity.Connection, error) {
	return c.repo.GetConnectionsByUserId(ctx, userId)
}

func NewConnectionService(repo repository.Connection, crypto Crypto) *ConnectionService {
	return &ConnectionService{repo: repo, crypto: crypto}
}
