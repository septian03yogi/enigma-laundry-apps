package config

import (
	"fmt"
	"os"

	"github.com/septian03yogi/enigmalaundryinc/utils/common"
)

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type Config struct {
	DbConfig
}

// env / getenv untuk membuat wadah agar pengaturan config(hos,port,name dll) fleksibel, menyesuaikan keadaan server
// Method
func (c *Config) ReadConfig() error {
	err := common.LoadEnv()
	if err != nil {
		return err
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Driver == "" {
		return fmt.Errorf("missing environment variables")
	}
	return nil
}

// constructor
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
