package config

import "github.com/solrac97gr/go-template/pkg/validator"

type Config struct {
	val *validator.Validator
	env *EnvironmentVariables
}

type EnvironmentVariables struct {
	DbUrl string `json:"db_url"`
}

func NewConfig(v *validator.Validator) *Config {
	return &Config{
		val: v,
		env: &EnvironmentVariables{},
	}
}

func (c *Config) GetEnvironmentVariables() *EnvironmentVariables {
	return c.env
}
