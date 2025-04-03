package mysqldb

import (
	"database/sql"
	"errors"
	"god/job/schedule"
	"god/pkg/config"
	"god/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func NewMysqlConnection() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	var sqlDB *sql.DB

	opts := config.MySQLConfig{}
	opts.LoadEnvs()
	dsn := opts.BuildConnection()

	logger.InfoF("Connect to database [%s] ...", opts.IP)

	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{Logger: gormLogger.Default.LogMode(gormLogger.Silent)}); err != nil {
		return nil, err
	}

	if sqlDB, err = db.DB(); err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(opts.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(opts.MaxLifeTime)
	sqlDB.SetMaxIdleConns(opts.MaxIdleConn)

	if err = sqlDB.Ping(); err != nil {
		return nil, errors.New("cannot ping mysql database")
	}

	healthCheckCron(sqlDB, opts.IP)
	return db, err
}

func healthCheckCron(sqlDB *sql.DB, ip string) {
	_, _ = schedule.RegisterScheduler("@every 0h5m0s", func() {
		var status = "OK"

		if err := sqlDB.Ping(); err != nil {
			status = err.Error()
		}

		logger.DebugF("try to ping to mysql=[%s] with msg=[%s]", ip, status)
	})
}

func Close(db *gorm.DB) error {
	var err error
	var sqlDB *sql.DB

	if sqlDB, err = db.DB(); err != nil {
		return err
	}

	err = sqlDB.Close()
	return err
}
