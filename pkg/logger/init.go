package logger

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"time"
)

func getTimeLogger() string {
	return time.Now().UTC().Add(7 * time.Hour).Format(`2006-01-02 15:04:05`)
}

func InfoF(format string, args ...any) {
	log.Info().
		Str("created_at", getTimeLogger()).
		Str("log_msg", fmt.Sprintf(format, args...)).Send()
}

func Info(msg string) {
	log.Info().
		Str("created_at", getTimeLogger()).
		Str("log_msg", msg).Send()
}

func InfoFields(msg map[string]interface{}) {
	log.Info().
		Str("created_at", getTimeLogger()).
		Fields(msg).Send()
}

func DebugF(format string, args ...any) {
	log.Debug().
		Str("created_at", getTimeLogger()).
		Str("log_msg", fmt.Sprintf(format, args...)).Send()
}

func Debug(msg string) {
	log.Debug().
		Str("created_at", getTimeLogger()).
		Str("log_msg", msg).Send()
}

func Error(msg string) {
	log.Debug().
		Str("created_at", getTimeLogger()).
		Str("log_msg", msg).Send()
}

func ErrorF(format string, args ...any) {
	log.Error().
		Str("created_at", getTimeLogger()).
		Str("log_msg", fmt.Sprintf(format, args...)).Send()
}
