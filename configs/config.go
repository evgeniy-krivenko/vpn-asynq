package configs

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/evgeniy-krivenko/vpn-asynq/internal/logger"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"os"
)

const (
	serviceName = "vpn-asynq"
	DbUsername  = "DB_USERNAME"
	DbPassword  = "DB_PASSWORD"
	DbHost      = "DB_HOST"
	DbPort      = "DB_PORT"
	DbName      = "DB_NAME"
	SslMode     = "SSL_MODE"
	SecretKey   = "SECRET_KEY"
	HttpPort    = "HTTP_PORT"
)

type Config struct {
	DbUsername string `json:"db_username"`
	DbPassword string `json:"db_password"`
	DbHost     string `json:"db_host"`
	DbPort     int    `json:"db_port"`
	DbName     string `json:"db_name"`
	SslMode    string `json:"ssl_mode"`
	SecretKey  string `json:"secret_key"`
	HttpPort   int    `json:"http_port"`
}

func InitConsulConfig(log logger.Logger, port int) error {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		return err
	}

	serviceID := serviceName
	address := getHostname()

	_ = port

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Port:    50051,
		Address: address,
		Check: &consulapi.AgentServiceCheck{
			GRPC:       fmt.Sprintf("%s:%d", address, 50051),
			GRPCUseTLS: false,
			Interval:   "30s",
			Timeout:    "30s",
		},
	}

	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		return err
	}
	log.Infof("successfully register service: %s:%v", address, port)

	kv := consul.KV()
	value, _, err := kv.Get(serviceName, nil)

	var cnf Config
	err = json.Unmarshal(value.Value, &cnf)
	if err != nil {
		return err
	}

	viper.Set(DbUsername, cnf.DbUsername)
	viper.Set(DbPassword, cnf.DbPassword)
	viper.Set(DbName, cnf.DbName)
	viper.Set(DbHost, cnf.DbHost)
	viper.Set(DbPort, cnf.DbPort)
	viper.Set(SslMode, cnf.SslMode)
	viper.Set(SecretKey, cnf.SecretKey)
	viper.Set(HttpPort, cnf.HttpPort)

	services, err := consul.Agent().ServicesWithFilter(`"ss" in Tags`)
	if err != nil {
		return err
	}
	for k, v := range services {
		fmt.Println("service ", k)
		fmt.Println("address", v.Address)
	}
	return nil
}

func MustPort() int {
	port := flag.Int(
		"http-port",
		0,
		"port for start checkup http",
	)

	flag.Parse()

	if *port == 0 {
		panic("not port")
	}

	return *port
}

func getHostname() (hostname string) {
	hostname, _ = os.Hostname()
	return
}
