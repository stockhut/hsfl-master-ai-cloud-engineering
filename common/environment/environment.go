package environment

import (
	"log"
	"os"
)

// GetRequiredEnvVar returns the specified variable value or exits the process with exit code 1
func GetRequiredEnvVar(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("No %s configured\n", key)
	}

	return value
}
