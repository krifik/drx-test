package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func NewConfiguration(filenames ...string) Config {
	godotenv.Load(filenames...)
	return &configImpl{}
}
