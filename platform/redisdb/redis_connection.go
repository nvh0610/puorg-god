package redisdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"god/job/schedule"
	"god/pkg/config"
	"god/pkg/logger"
	"time"
)

func NewRedisConnection() (*redis.Client, error) {
	opt := config.RedisConfig{}
	opt.LoadEnvs()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", opt.Host, opt.Port),
		Password: opt.Password,
		DB:       opt.DB,
		Username: opt.UserName,
	})

	status, err := rdb.Ping(ctx).Result()
	logger.InfoF("try to ping to redis connection=[%s:%s] with status=[%s]", opt.Host, opt.Port, status)
	if status != "PONG" {
		return nil, errors.New("redis PING failed")
	}

	healthCheckCron(rdb, []string{fmt.Sprintf("%s:%s", opt.Host, opt.Port)})

	return rdb, err
}

func healthCheckCron(rdb redis.UniversalClient, addrs []string) {
	_, _ = schedule.RegisterScheduler("@every 0h5m0s", func() {
		status, _ := rdb.Ping(context.Background()).Result()
		logger.InfoF("try to ping to redis=[%v] with status=[%s]", addrs, status)
	})
}
