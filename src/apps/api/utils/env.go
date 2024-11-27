package utils

import "os"

func IsAPIInProdMode() bool {
	return os.Getenv("API_MODE") == "prod"
}
