package service

import (
	"context"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/entity"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/repository"
)

// define interfaces for our services

type Connection interface {
	GetConnections(ctx context.Context, userId int64) ([]entity.Connection, error)
	CreateConnection(ctx context.Context, usr *entity.User, serverId int) (int, error)
}

type Service struct {
	Connection
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Connection: NewConnectionService(repos.Connection, new(CryptoService)),
	}
}
