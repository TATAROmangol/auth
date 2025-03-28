package postgres

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func LoadConfig() (Config, error){
	host, exist := os.LookupEnv("PG_HOST")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_HOST")
	}

	sPort, exist := os.LookupEnv("PG_PORT")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_PORT")
	}
	port, err := strconv.Atoi(sPort)
	if err != nil{
		return Config{}, fmt.Errorf("failed load PG_PORT: %v", err)
	}

	user, exist := os.LookupEnv("PG_USER")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_USER")
	}

	password, exist := os.LookupEnv("PG_PASSWORD")
	if !exist {
		return Config{}, fmt.Errorf("failed load PG_PASSWORD")
	}

	dbName, exist := os.LookupEnv("DB_NAME")
	if !exist {
		return Config{}, fmt.Errorf("failed load DB_NAME")
	}

	return Config{
		host,
		port,
		user,
		password,
		dbName,
	}, nil
}

func (cfg Config) GetConnectPath() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
}
