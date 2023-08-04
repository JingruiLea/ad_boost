package envs

import "os"

func IsDev() bool {
	env := os.Getenv("ENV")
	if env == "" {
		return true
	}
	if env == "dev" {
		return true
	}
	return false
}
