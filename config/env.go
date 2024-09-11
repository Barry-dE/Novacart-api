// Package config provides functionality to load and manage application configuration settings.
package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

// Configuration struct holds the configuration settings for the application
type Configuration struct {
	Port       string
	PublicHost string
	DBPassword string
	DBAdrress  string
	DBUser     string
	DBName     string
}

// Envs is a global variable that holds the application's configuration settings
var Envs = initConfig()

// initConfig initializes the configuration by loading environment variables and setting defaults
func initConfig() Configuration {
	// Load environment variables from a .env file if it exists
	godotenv.Load()

	// Return a Configuration struct with values from environment variables or defaults if ".env" does not exist
	return Configuration{
		PublicHost: getEnvironment("PUBLIC_HOST", "http://localhost"),
		Port:       getEnvironment("PORT", "8080"),
		DBName:     getEnvironment("DB_NAME", "novacart"),
		DBUser:     getEnvironment("DB_USER", "root"),
		DBPassword: getEnvironment("DB_PASSWORD", "devopsoncloud"),
		DBAdrress:  fmt.Sprintf("%s:%s", getEnvironment("DB_HOST", "127.0.0.1"), getEnvironment("DB_PORT", "3306")),
	}
}

// getEnvironment retrieves the value of an environment variable or returns a fallback value if the variable is not set
func getEnvironment(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return fallback
}
