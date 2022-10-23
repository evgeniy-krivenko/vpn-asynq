package conn_grpc

import (
	"context"
	pb "github.com/evgeniy-krivenko/vpn-api/api"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/logger"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ConnectionTransport struct {
	service.Connection
	log logger.Logger
	pb.UnimplementedConnectionServiceServer
}

func NewConnectionTransport(srv service.Connection, log logger.Logger) *ConnectionTransport {
	return &ConnectionTransport{Connection: srv, log: log}
}

func (c *ConnectionTransport) GetConnections(ctx context.Context, req *pb.GetConnectionsReq) (*pb.GetConnectionsRes, error) {
	var resp pb.GetConnectionsRes
	conns, err := c.Connection.GetConnections(ctx, req.GetUserId())
	if err != nil {
		c.log.WithContextReqId(ctx).
			Errorf("error to get connections for user: %d: w%", req.GetUserId(), err)
	}

	for _, conn := range conns {
		resp.Connections = append(resp.Connections,
			&pb.Connection{
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
func (c *ConnectionTransport) GetConnectionInfo(context.Context, *pb.GetConnectionInfoReq) (*pb.Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionInfo not implemented")
}
func (c *ConnectionTransport) GetServers(context.Context, *pb.GetServersReq) (*pb.GetServersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServers not implemented")
}
func (c *ConnectionTransport) GetConfig(context.Context, *pb.GetConfigReq) (*pb.GetConfigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (c *ConnectionTransport) ActivateConnection(context.Context, *pb.ActivateConnectionReq) (*pb.ActivateConnectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateConnection not implemented")
}
func (c *ConnectionTransport) DeactivateConnection(context.Context, *pb.DeactivateConnectionReq) (*pb.DeactivateConnectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateConnection not implemented")
}
