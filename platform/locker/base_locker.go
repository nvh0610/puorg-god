package locker

import "time"

type BaseLock interface {
	Lock(key string, TTL time.Duration) error
	NonIntervalLock(key string, TTL time.Duration) error
	UnLock(key string) error
}
