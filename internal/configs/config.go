package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

const (
	DevEnv  = "local"
	ProdEnv = "production"
)

type Config struct {
	Debug bool   `envconfig:"DEBUG"`
	Env   string `envconfig:"ENV"`

	DB struct {
		DatabaseUrl string `envconfig:"DATABASE_URL"`
	}

	HTTPConfig struct {
		Port               string        `yaml:"port"`
		ReadTimeout        time.Duration `yaml:"readTimeout"`
		WriteTimeout       time.Duration `yaml:"writeTimeout"`
		MaxHeaderMegabytes int           `yaml:"maxHeaderBytes"`
	} `yaml:"http"`

	Auth struct {
		AccessTokenTTL string `yaml:"accessTokenTTL"`
	}

	Limiter struct {
		RPS   int           `yaml:"rps"`
		Burst int           `yaml:"burst"`
		TTL   time.Duration `yaml:"ttl"`
	} `yaml:"limiter"`
}

func Init(configsDir string) (Config, error) {
	var cfg Config

	readEnv(&cfg)
	readFile(&cfg, configsDir)

	return cfg, nil
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config, configsDir string) {
	f, err := os.Open(configsDir + "/main.yml")
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := godotenv.Load()
	if err != nil {
		// processError(err)
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
