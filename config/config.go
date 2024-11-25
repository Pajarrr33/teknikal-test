package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/joho/godotenv"
	"github.com/golang-jwt/jwt/v5"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
}

type ApiConfig struct {
	ApiPort string
}

type JwtConfig struct {
	Issuer 		string `json:"issuer"`
	SecretKey   []byte `json:"secretKey"`
	Method      *jwt.SigningMethodHMAC
	Expire      time.Duration
}

type Config struct {
	DbConfig
	ApiConfig
	JwtConfig
}

func (c *Config) Init() error{
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}
	tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))

	c.JwtConfig = JwtConfig{
		Issuer: 		os.Getenv("TOKEN_ISSUE"),
		SecretKey:      []byte(os.Getenv("TOKEN_SECRET")),
		Method:         jwt.SigningMethodHS256,
		Expire:         time.Duration(tokenExpire) * time.Minute,
	}

	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" || c.DbConfig.Name == "" || c.DbConfig.Driver == "" || c.ApiConfig.ApiPort == "" ||
		c.JwtConfig.Issuer == "" || c.JwtConfig.Expire < 0 || len(c.SecretKey) == 0 {
		return fmt.Errorf("missing required environment")
	}

	return nil
}

func GetConfig() (*Config,error) {
	config := &Config{}
	err := config.Init()
	if err != nil {
		return nil, err
	}
	return config,nil
}