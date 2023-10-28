package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ProgramConfig struct {
	ServerPort int
	DBPort     int
	DBHost     string
	DBUser     string
	DBPass     string
	DBName     string
	Secret     string
}

func InitConfig() *ProgramConfig {
	var res = new(ProgramConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cant start program failed to load configuration")
	}

	return res
}

func loadConfig() *ProgramConfig {
	var res = new(ProgramConfig)

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Config : Cannot load config file,", err.Error())
		return nil
	}

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Fatal("Config : invalid port value", err.Error())
		}
		res.ServerPort = port
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Fatal("Config : invalid port value", err.Error())
		}
		res.DBPort = port
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = val
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	}
	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}

	return res
}
