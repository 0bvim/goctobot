package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	PersonalGithubToken string `env:"PERSONAL_GITHUB_TOKEN,required,notEmpty"`
}

func LoadConfig[T any](cfg *T) error {
	if cfg == nil {
		return fmt.Errorf("config pointer cannot be nil")
	}

	v := reflect.ValueOf(cfg)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("config must be a pointer to a struct")
	}

	if err := env.Parse(cfg); err != nil {
		errors := strings.ReplaceAll(err.Error(), "; ", "\n")
		errors = strings.TrimPrefix(errors, "env: ")

		return fmt.Errorf("%s", errors)
	}

	return nil
}
