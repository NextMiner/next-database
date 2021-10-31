package next_database

import (
	"crypto/tls"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type DatabaseConfig struct {
	Network   string            `json:"network"`
	Host      string            `json:"host"`
	Port      int               `json:"port"`
	Password  string            `json:"password"`
	DB        int               `json:"db"`
	TLS       *TLSClientOptions `json:"tls"`
}

func (dc *DatabaseConfig) Addr() string {
	return dc.Host + ":" + strconv.Itoa(dc.Port)
}

func (dc *DatabaseConfig) ToRedisOptions() *redis.Options {
	var tlsConfig *tls.Config

	if dc.TLS != nil {
		tlsConfig = dc.TLS.ToTLSConfig()
	}

	return &redis.Options{
		Network:   dc.Network,
		Addr:      dc.Addr(),
		Password:  dc.Password,
		DB:        dc.DB,
		TLSConfig: tlsConfig,
	}
}