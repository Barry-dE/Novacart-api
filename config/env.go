package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

type Configuration struct {
	Port       string
	PublicHost string

	DBPassword string
	DBAdrress  string
	DBUser     string
	DBName     string
}

var Envs = initConfig()


func initConfig() Configuration {
godotenv.Load()
	return Configuration{
		PublicHost: getEnvironment("PUBLIC_HOST", "http//localhost"),
		Port: getEnvironment("PORT", "8080"),
		DBName: getEnvironment("DB_NAME", "novacart"),
		DBUser: getEnvironment("DB_USER", "root"),
		DBPassword: getEnvironment("DB_PASSWORD", "devopsoncloud"),
		DBAdrress: fmt.Sprintf("%s:%s", getEnvironment("DB_HOST", "127.0.0.1"), getEnvironment("DB_PORT", "5173")),

	}
}


func getEnvironment(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok{
		return value
	}

	return fallback
}



