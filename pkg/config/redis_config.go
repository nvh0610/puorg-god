package config

import "fmt"

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	UserName string
	DB       int
}

func (c *RedisConfig) BuildConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%d", c.UserName, c.Password, c.Host, c.DB)
}

func (c *RedisConfig) LoadEnvs() {
	c.Host = StringEnv(`REDIS_HOST`)
	c.Port = StringEnv(`REDIS_PORT`)
	c.Password = StringEnv(`REDIS_PASSWORD`)
	c.UserName = StringEnv(`REDIS_USERNAME`)
	c.DB = IntEnvF("REDIS_DB", 0)
}
