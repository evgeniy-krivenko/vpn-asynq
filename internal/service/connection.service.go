package service

import (
	"context"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/entity"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/repository"
)

type ConnectionService struct {
	repo repository.Connection
}

func (c ConnectionService) GetConnections(ctx context.Context, userId int64) ([]entity.Connection, error) {
	return c.repo.GetConnectionsByUserId(ctx, userId)
}

func NewConnectionService(repo repository.Connection) *ConnectionService {
	return &ConnectionService{repo: repo}
}
