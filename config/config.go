package config

import (
	"github.com/Netflix/go-env"
	"github.com/rs/zerolog/log"
)

type Environment struct {
	AppName                     string `env:"APP_NAME" mapstructure:"APP_NAME"`
	AppEnv                      AppEnv `env:"APP_ENV" mapstructure:"APP_ENV"`
	AppPort                     int    `env:"APP_PORT" mapstructure:"APP_PORT"`
	SystemShutdownTimeOutSecond int    `env:"APP_SYSTEM_SHUTDOWN_TIMEOUT_SECOND" mapstructure:"APP_SYSTEM_SHUTDOWN_TIMEOUT_SECOND"`
	AllowOrigins                string `env:"CORS_ALLOW_ORIGINS" mapstructure:"CORS_ALLOW_ORIGINS"`

	JwtEnv `mapstructure:",squash"`

	MongoEnv `mapstructure:",squash"`

	AwsEnv `mapstructure:",squash"`

	S3Env `mapstructure:",squash"`
}

type AwsEnv struct {
	AwsAccessKey string `env:"AWS_ACCESS_KEY" mapstructure:"AWS_ACCESS_KEY"`
	AwsSecret    string `env:"AWS_SECRET_KEY" mapstructure:"AWS_SECRET_KEY"`
	AwsRegion    string `env:"AWS_REGION" mapstructure:"AWS_REGION"`
}

type S3Env struct {
	S3Bucket              string `env:"S3_BUCKET_ID" mapstructure:"S3_BUCKET_ID"`
	S3PreSignDurationHour int    `env:"S3_PRE_SIGN_DURATION_HOUR" mapstructure:"S3_PRE_SIGN_DURATION_HOUR"`
}

type MongoEnv struct {
	MongoDbHost              string `env:"MongoDB_HOST" mapstructure:"MongoDB_HOST"`
	MongoDbPort              string `env:"MongoDB_PORT" mapstructure:"MongoDB_PORT"`
	MongoDbDatabase          string `env:"MongoDB_NAME" mapstructure:"MongoDB_NAME"`
	MongoDbUsername          string `env:"MongoDB_USER" mapstructure:"MongoDB_USER"`
	MongoDbPassword          string `env:"MongoDB_PASS" mapstructure:"MongoDB_PASS"`
	MongoDbTimeout           int    `env:"MongoDB_TIMEOUT" mapstructure:"MongoDB_TIMEOUT"`
	MongoDbOption            string `env:"MongoDB_OPTION" mapstructure:"MongoDB_OPTION"`
	MongoDbMaxPoolSize       int    `env:"MongoDB_MAX_POOL_SIZE" mapstructure:"MongoDB_MAX_POOL_SIZE"`
	MongoDbMaxConnectionSize int    `env:"MongoDB_MAX_CONNECTION_SIZE" mapstructure:"MongoDB_MAX_CONNECTION_SIZE"`
}

type JwtEnv struct {
	JWTSecret         string `env:"JWT_SECRET" mapstructure:"JWT_SECRET"`
	JWTRefreshSecret  string `env:"JWT_REFRESH_SECRET" mapstructure:"JWT_REFRESH_SECRET"`
	JWTExpires        int    `env:"JWT_EXPIRES" mapstructure:"JWT_EXPIRES"`
	JWTRefreshExpires int    `env:"JWT_REFRESH_EXPIRES" mapstructure:"JWT_REFRESH_EXPIRES"`
	JWTCookieName     string `env:"JWT_COOKIE_NAME" mapstructure:"JWT_COOKIE_NAME"`
}

type AppEnv string

func (e AppEnv) IsLocal() bool {
	return e == "localhost"
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal().Msgf("can't load env: %v", err)
		return nil, err
	}
	return &environment, nil
}
