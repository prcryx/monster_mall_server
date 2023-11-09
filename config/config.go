package config

import (
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/prcryx/monster_mall/internal/common/exception"
)

type EnvConfig struct {
	ServerPort         uint16 `config:"port" mapstructure:"SERVER_PORT"`
	ServerName   string `config:"server_name" mapstructure:"SERVER_NAME"`
	ServerAddr   string `config:"server_addr" mapstructure:"SERVER_ADDR"`
	SecretKey    string `config:"secret_key" mapstructure:"SECRET_KEY"`
	TokenExp     string `config:"token_exp" mapstructure:"TOKEN_EXPIRY_IN_MIN"`
	PostgresUser string `config:"postgres_user" mapstructure:"PG_USER"`
	PostgresPass string `config:"postgres_pass" mapstructure:"PG_PASS"`
	PostgresDb   string `config:"postgres_db" mapstructure:"PG_DB"`
	PostgresSSL  bool   `config:"postgres_ssl" mapstructure:"PG_SSL"`
	PostgresPort uint16 `config:"postgres_port" mapstructure:"PG_PORT"`
	RedisPort    uint16 `config:"redis_port" mapstructure:"REDIS_PORT"`
}

func LoadConfig() (*EnvConfig, error) {
	var config = new(EnvConfig)
	configMap, err := godotenv.Read()
	if err != nil {
		return nil, exception.InternalServerError(err.Error())
	}

	decodeConfigMap(configMap, config)

	return config, nil
}

func decodeConfigMap(in map[string]string, out *EnvConfig) {
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			ZeroFields:       true,
			WeaklyTypedInput: true,
			TagName:          "mapstructure",
			Result:           out,
		},
	)

	decoder.Decode(in)

	if err != nil {
		return
	}
}
