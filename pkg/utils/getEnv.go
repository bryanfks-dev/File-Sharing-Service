package utils

import "os"

// GetEnv is a helper function that returns the value of an environment
// variable or a default value if the environment variable is not set.
//
// key - the name of the environment variable
// defaultValue - the default value to return if the environment variable is not set
func GetEnv(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)

	// If the environment variable is not set, return the default value
	if !ok {
		return defaultValue
	}

	return val
}
