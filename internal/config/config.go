package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type EnvType string

func (e EnvType) String() string {
	return string(e)
}

const (
	EnvTypeDev EnvType = "development"
)

// Struct configuration struct.
type Struct struct {
	GrpcPort string `yaml:"grpcPort" env:"GRPCPORT" env-default:"50052"`
	HttpPort string `yaml:"httpPort" env:"HTTPPORT" env-default:"localhost:50052"`
	LogLevel int    `yaml:"logLevel" env:"LOGLEVEL" env-default:"1"`
	Env      string `yaml:"env" env:"ENV" env-default:"development"`
}

// Config configuration.
var Config Struct

func Init() {
	filePath := "config.yml"
	if os.Getenv("ENV") == "local" {
		filePath = "config_local.yml"
	}
	err := cleanenv.ReadConfig(filePath, &Config)

	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

}
