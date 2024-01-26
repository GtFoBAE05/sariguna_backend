package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type AppConfig struct {
	SERVERPORT int
	DBPORT     int
	DBHOST     string
	DBUSER     string
	DBPASS     string
	DBNAME     string
}

func ConnectPostgres(cfg *AppConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHOST, cfg.DBPORT, cfg.DBUSER, cfg.DBPASS, cfg.DBNAME,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func InitConfig() *AppConfig {
	var res = new(AppConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *AppConfig {
	var res = new(AppConfig)

	godotenv.Load(".env")

	if val, found := os.LookupEnv("SERVERPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid port value,", err.Error())
			return nil
		}
		res.SERVERPORT = port
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid db port value,", err.Error())
			return nil
		}
		res.DBPORT = port
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHOST = val
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUSER = val
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPASS = val
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBNAME = val
	}

	return res

}
