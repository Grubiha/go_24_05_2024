package env

import (
	"errors"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var errEnvNotFound = errors.New("file not found: \".env\"")

func Load() {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		slog.Error(err.Error())
		panic(errEnvNotFound)
	} else if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	err = godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
