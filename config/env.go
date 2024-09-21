package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

/*
Config struct contains the configuration of the server
*/
type Config struct {
	PublicHost string
	Port       string
	DBAddr     string
	DBUser     string
	DBPass     string
	DBName     string
}

/*
initConfig function initializes the configuration of the server
*/
func initConfig() Config {

	// load The Environment Variables
	godotenv.Load()

	return Config{
		PublicHost: GetEnv("PUBLIC_HOST", "http://localhost"),
		Port:       GetEnv("PORT", "8080"),
		DBUser:     GetEnv("DB_USER", "root"),
		DBPass:     GetEnv("DB_PASS", "root"),
		DBAddr:     fmt.Sprintf("%s:%s", GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "3306")),
		DBName:     GetEnv("DB_NAME", "ecom"),
	}
}

var Envs = initConfig()

/*
GetEnv function returns the value of the environment variable if it exists, otherwise it returns the fallback value
*/
func GetEnv(key, fallback string) string {
	// Get the value of the environment variable
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
