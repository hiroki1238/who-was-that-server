package config

import (
	"os"
	"strconv"
)

type Config struct {
	Database DatabaseConfig
	Redis    RedisConfig
	Cognito  CognitoConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type CognitoConfig struct {
	UserPoolID       string
	ClientID         string
	IdentityPoolID   string
	Region           string
}

type ServerConfig struct {
	Port     int
	JWTSecret string
}

func Load() (*Config, error) {
	cfg := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "who_was_that_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Cognito: CognitoConfig{
			UserPoolID:     getEnv("COGNITO_USER_POOL_ID", ""),
			ClientID:       getEnv("COGNITO_CLIENT_ID", ""),
			IdentityPoolID: getEnv("COGNITO_IDENTITY_POOL_ID", ""),
			Region:         getEnv("AWS_REGION", "ap-northeast-1"),
		},
		Server: ServerConfig{
			Port:      getEnvAsInt("SERVER_PORT", 8080),
			JWTSecret: getEnv("JWT_SECRET", "default_secret"),
		},
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	strValue := os.Getenv(key)
	if strValue == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return defaultValue
	}

	return intValue
}
