package locker

import (
	"context"
	"crawler_go_facebook/pkg/logger"
	"crawler_go_facebook/platform/redisdb"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	// ErrNotObtained is returned when a lock_manager cannot be obtained.
	ErrNotObtained = errors.New("redislock: not obtained")

	// ErrLockNotHeld is returned when trying to release an inactive lock_manager.
	ErrLockNotHeld = errors.New("redislock: lock_manager not held")
)

type RedisLocker struct {
	client *redisdb.RedisConnection
}

func NewRedisLocker(client *redisdb.RedisConnection) *RedisLocker {
	return &RedisLocker{
		client: client,
	}
}

func (locker *RedisLocker) Lock(key string, TTL time.Duration) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), TTL)
	defer cancel()

	chDone := make(chan int)

	go func(ctx context.Context) {
		defer close(chDone)
		var nx bool

		for {
			if nx, err = locker.client.DB.SetNX(ctx, key, 1, TTL).Result(); err != nil {
				chDone <- 1
			}

			if nx {
				chDone <- 1
				break
			}
		}
	}(ctx)

	select {
	case <-ctx.Done():
		return ErrNotObtained
	case <-chDone:
		return err
	}
}

func (locker *RedisLocker) NonIntervalLock(key string, TTL time.Duration) error {
	var nx bool
	var err error

	if nx, err = locker.client.DB.SetNX(context.Background(), key, 1, TTL).Result(); err != nil {
		return err
	}

	if !nx {
		return ErrNotObtained
	}
	logger.InfoF("lock key=[%s] success %v", key, err)
	return err
}

func (locker *RedisLocker) UnLock(key string) error {
	var deleteScript = redis.NewScript(`
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`)
	_, err := deleteScript.Run(context.Background(), locker.client.DB, []string{key}, 1).Result()
	if err == nil {
		logger.InfoF("unlock key=[%s] success", key)
	}
	return err
}
