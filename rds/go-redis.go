package rds

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RDSInfo struct {
	Host string
	Port int
	Pass string
}

func Conn(rdsi *RDSInfo) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rdsi.Host, rdsi.Port),
		Password: rdsi.Pass,
		DB:       0,
	})
}
