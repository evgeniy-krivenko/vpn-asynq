package conn_grpc

import (
	"context"
	"github.com/evgeniy-krivenko/vpn-api/gen/conn_service"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/logger"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ConnectionTransport struct {
	service.Connection
	log logger.Logger
	conn_service.UnimplementedConnectionServiceServer
}

func NewConnectionTransport(srv service.Connection, log logger.Logger) *ConnectionTransport {
	return &ConnectionTransport{Connection: srv, log: log}
}

func (c *ConnectionTransport) GetConnections(ctx context.Context, req *conn_service.GetConnectionsReq) (*conn_service.GetConnectionsRes, error) {
	var resp conn_service.GetConnectionsRes
	conns, err := c.Connection.GetConnections(ctx, req.GetUserId())
	if err != nil {
		c.log.WithContextReqId(ctx).
			Errorf("error to get connections for user: %d: w%", req.GetUserId(), err)
	}
	resp.Connections = make([]*conn_service.Connection, len(conns))

	for _, conn := range conns {
		resp.Connections = append(resp.Connections,
			&conn_service.Connection{
				Id:           int64(conn.Id),
				Location:     conn.Location,
				Port:         uint64(conn.Port),
				UserId:       int64(conn.UserId),
				IpAddress:    conn.IpAddress,
				ServerId:     int64(conn.ServerId),
				IsActive:     conn.IsActive,
				LastActivate: timestamppb.New(conn.LastActivate.Time),
			})
	}

	return &resp, nil
}
func (c *ConnectionTransport) GetConnectionInfo(context.Context, *conn_service.GetConnectionInfoReq) (*conn_service.Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionInfo not implemented")
}
func (c *ConnectionTransport) GetServers(context.Context, *conn_service.GetServersReq) (*conn_service.GetServersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServers not implemented")
}

func (c *ConnectionTransport) CreateConnection(context.Context, *conn_service.CreateConnectionReq) (*conn_service.CreateConnectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateConnection not implemented")
}

func (c *ConnectionTransport) GetConfig(context.Context, *conn_service.GetConfigReq) (*conn_service.GetConfigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (c *ConnectionTransport) ActivateConnection(context.Context, *conn_service.SwitchConnectionReq) (*conn_service.SwitchConnectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateConnection not implemented")
}
func (c *ConnectionTransport) DeactivateConnection(context.Context, *conn_service.SwitchConnectionReq) (*conn_service.SwitchConnectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateConnection not implemented")
}
