package utils

import (
	"log"
	"os"
	"runtime/debug"

	"github.com/joho/godotenv"
)

func init() {
	// Load the .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		log.Printf("An error occurred while trying to load the .env file: %v", err)
	}
}

func GetEnvVar(name string) string {
	envVar := os.Getenv(name)
	if envVar == "" {
		debug.PrintStack()
		log.Printf("env var %v must be set", name)
	}
	return envVar
}
