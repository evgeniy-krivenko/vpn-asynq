package main

import (
	"context"
	"fmt"
	"github.com/evgeniy-krivenko/vpn-api/gen/conn_service"
	"github.com/evgeniy-krivenko/vpn-asynq/configs"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/logger"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/repository"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/service"
	conn_grpc "github.com/evgeniy-krivenko/vpn-asynq/internal/transport/grpc/v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healphpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"net/http"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	lgr := logrus.New()
	lgr.SetFormatter(new(logrus.JSONFormatter))

	log := logger.NewLogrusLogger(lgr)

	viper.AutomaticEnv()
	httpPort := configs.MustPort()
	if err := configs.InitConsulConfig(log, httpPort); err != nil {
		log.Fatalf("%s", err.Error())
	}

	cfg := repository.Config{
		Host:     viper.GetString(configs.DbHost),
		Port:     viper.GetString(configs.DbPort),
		Username: viper.GetString(configs.DbUsername),
		Password: viper.GetString(configs.DbPassword),
		DBName:   viper.GetString(configs.DbName),
		SSLMode:  viper.GetString(configs.SslMode),
	}

	db, err := repository.NewPostgresDB(ctx, cfg)
	if err != nil {
		log.Fatalf("error load database: %s", err.Error())
	}
	defer db.Close()

	log.Infof("successful connected to db by host: %s and port: %s", cfg.Host, cfg.Port)

	go func() {
		http.HandleFunc("/check", check)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil); err != nil {
			log.Fatalf("error to start checked service: %s", err.Error())
		}
	}()

	/*
			ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*10)
			defer cancel()
		в цикле создаем коннекты к разным ss сервиса по ключам из консула
		добавляем в мапу по ключу сервиса.
			conn, err := grpc.DialContext(ctxWithTimeout, "localhost:50051",
				grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
			if err != nil {
				log.Fatalf("error connect to grpc server %s", err.Error())
			}
			defer conn.Close()

			cs := conn_service.NewConnectionServiceClient(conn)
	*/

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error to open connection: %w", err)
	}

	repo := repository.NewRepository(db, log)
	services := service.NewService(repo)

	grpcServer := grpc.NewServer()

	conn_service.RegisterConnectionServiceServer(
		grpcServer,
		conn_grpc.NewConnectionTransport(services, log),
	)

	healthGrpcServer := health.NewServer()
	healphpb.RegisterHealthServer(grpcServer, healthGrpcServer)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("error to start server: %w", err)
		}
	}()

	log.Infof("successful starting server on port %s", ":50051")
	select {
	case <-ctx.Done():
		log.Info("graceful stop server")
		grpcServer.GracefulStop()
	}
}

func check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Consul check")
}
