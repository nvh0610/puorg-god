package config

import (
	"fmt"
	"time"
)

type MySQLConfig struct {
	Username    string
	Password    string
	IP          string
	DB          string
	MaxIdleConn int
	MaxOpenConn int
	MaxLifeTime time.Duration
}

func (*MySQLConfig) IPEnv() string {
	host := StringEnv(`DATABASE__HOST`)
	port := StringEnv(`DATABASE__PORT`)
	return fmt.Sprintf("%s:%s", host, port)
}

func (*MySQLConfig) UsernameEnv() string {
	return StringEnv(`DATABASE__USER`)
}

func (*MySQLConfig) PasswordEnv() string {
	return StringEnv(`DATABASE__PASSWORD`)
}

func (*MySQLConfig) DatabaseEnv() string {
	return StringEnv(`DATABASE__NAME`)
}

func (c *MySQLConfig) LoadUriEnv() {
	uri := StringEnv(`DATABASE_MYSQL_URI`)
	v := ParseConnectionString(uri)
	c.Username = v.User.Username()
	c.Password, _ = v.User.Password()
	c.IP = v.Host
	c.DB = v.DatabaseName
	c.MaxIdleConn = IntEnvF("DB_MAX_IDLE_CONNECTIONS", 10)
	c.MaxOpenConn = IntEnvF("DB_MAX_CONNECTIONS", 100)
	c.MaxLifeTime = time.Duration(IntEnvF("DB_MAX_LIFETIME_CONNECTIONS", 1)) * time.Hour
}

func (c *MySQLConfig) LoadEnvs() {
	c.IP = c.IPEnv()
	c.Username = c.UsernameEnv()
	c.Password = c.PasswordEnv()
	c.DB = c.DatabaseEnv()
	c.MaxIdleConn = IntEnvF("DB_MAX_IDLE_CONNECTIONS", 10)
	c.MaxOpenConn = IntEnvF("DB_MAX_CONNECTIONS", 100)
	c.MaxLifeTime = time.Duration(IntEnvF("DB_MAX_LIFETIME_CONNECTIONS", 1)) * time.Hour
}

func (c *MySQLConfig) BuildConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.IP, c.DB)
}
