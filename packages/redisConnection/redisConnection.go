package redisconnection

import (
	"crypto/tls"
	"os"

	"github.com/redis/go-redis/v9"
)

func RedisConnection() *redis.ClusterClient {
	redis_uri := os.Getenv("REDIS_URI")
	redis_password := os.Getenv("REDIS_PASSWORD")
	redis_tls := os.Getenv("REDIS_TLS")
	redis_username := os.Getenv("REDIS_USERNAME")

	redisOptions := &redis.ClusterOptions{
		Addrs:    []string{redis_uri},
		Username: redis_username,
		Password: redis_password,
	}

	if redis_tls == "true" {
		redisOptions.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	rdb := redis.NewClusterClient(redisOptions)
	return rdb
}
