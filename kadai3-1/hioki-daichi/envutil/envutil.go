package envutil

import (
	"os"
	"strconv"
)

// GetIntEnvOrElse returns numeric-converted value if an environment variable of the specified key exists,
// and if not, returns the specified default value.
func GetIntEnvOrElse(key string, defaultValue int) (int, error) {
	if s, ok := os.LookupEnv(key); ok {
		return strconv.Atoi(s)
	}
	return defaultValue, nil
}
