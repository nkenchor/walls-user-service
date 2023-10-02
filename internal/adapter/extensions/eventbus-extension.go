package extensions

import (
	redisClient "github.com/redis/go-redis/v9"
	"strings"
	redis "walls-user-service/internal/adapter/repository/redis"
	logger "walls-user-service/internal/core/helper/log-helper"
)

func StartEventBus(ebType string) *redisClient.Client {

	switch ebType {
	case strings.ToLower(ebType):
		logger.LogEvent("INFO", "Initializing Redis!")
		redisRepo := redis.ConnectToRedis()
		return redisRepo
	}
	return nil

}
