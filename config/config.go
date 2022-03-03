package config

import (
	"github.com/joho/godotenv"
	"os"
	"restfull-api-arcticles/helper"
)

type Config interface {
	Get(key string) string
}

type ConfigImpl struct {

}

func (config *ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filename ...string) Config  {
	err := godotenv.Load(filename...)
	helper.PanicIfError(err)
	return &ConfigImpl{}
}