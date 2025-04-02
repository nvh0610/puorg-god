package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"god/pkg/logger"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type ServerInformation struct {
	Domain       string
	Port         int
	DatabaseName string
	Fragments    url.Values
	Queries      url.Values
	url.URL
}

func LoadPathEnv(path string) error {
	err := godotenv.Load(path)
	return err
}

func StringEnvF(key string, dval string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return dval
	}
	return v
}

func StringEnv(key string) string {
	return os.Getenv(key)
}

func StringEnvArray(key string, delimiter string) []string {
	return strings.Split(StringEnv(key), delimiter)
}

func DomainStringEnv(key string) string {
	domain := os.Getenv(key)
	l := len(domain) - 1
	if strings.HasPrefix(domain, "/") && domain[l] == '/' {
		domain = domain[:l]
	}
	return domain
}

func IntEnvF(key string, dval int) int {
	num, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return dval
	}
	return num
}

func IntEnv(key string) int {
	num, _ := strconv.Atoi(os.Getenv(key))
	return num
}

func FloatEnv(key string) float64 {
	num, _ := strconv.ParseFloat(os.Getenv(key), 64)
	return num
}

func FloatEnvF(key string, dval float64) float64 {
	num, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return dval
	}
	return num
}

func BooleanEnv(key string) bool {
	num, _ := strconv.ParseBool(os.Getenv(key))
	return num
}

func BooleanEnvF(key string, dval bool) bool {
	num, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return dval
	}
	return num
}

func ParseConnectionString(s string) *ServerInformation {
	if len(s) == 0 {
		return nil
	}
	u, err := url.Parse(s)

	if err != nil {
		logger.InfoF("failed to parse url %s", err.Error())
		return nil
	}

	fragments, err := url.ParseQuery(u.Fragment)
	if err != nil {
		logger.InfoF("failed to parse Fragment %s", err.Error())
	}
	hostname, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		logger.InfoF("failed to parse Host %s", err.Error())
	}

	nport, err := strconv.Atoi(port)
	if err != nil {
		logger.InfoF("failed to parse Port %s", err.Error())
	}

	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		logger.InfoF("failed to parse RawQuery %s", err.Error())
	}

	v := &ServerInformation{
		hostname,
		nport,
		strings.Replace(u.Path, "/", "", -1),
		fragments,
		query,
		*u,
	}
	return v
}

func (s *ServerInformation) FormatAsGo() string {
	return fmt.Sprintf("")
}
