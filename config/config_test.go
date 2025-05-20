package config_test

import (
	"os"
	"testing"

	"github.com/0bvim/goctobot/config"
)

func Test_LoadConfig(t *testing.T) {
	t.Run("Basic environment variable loading", func(t *testing.T) {
		os.Setenv("ABC", "123")
		defer os.Unsetenv("ABC")

		type Config struct {
			PersonalGithubToken string `env:"ABC,required,notEmpty"`
		}

		var cfg Config

		err := config.LoadConfig(&cfg)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if cfg.PersonalGithubToken != "123" {
			t.Fatalf("expected PersonalGithubToken to be '123', got '%s'", cfg.PersonalGithubToken)
		}
	})

	t.Run("Required but missing environment variable", func(t *testing.T) {
		os.Unsetenv("ABC")

		type Config struct {
			PersonalGithubToken string `env:"ABC,required"`
		}

		var cfg Config

		err := config.LoadConfig(&cfg)
		if err == nil {
			t.Fatal("expected error for missing required env var, got nil")
		}
	})

	t.Run("Default value for missing environment variable", func(t *testing.T) {
		os.Unsetenv("XYZ")

		type Config struct {
			ApiKey string `env:"XYZ" envDefault:"default_key"`
		}

		var cfg Config

		err := config.LoadConfig(&cfg)
		if err != nil {
			t.Fatalf("expected no error with default value, got %v", err)
		}

		if cfg.ApiKey != "default_key" {
			t.Fatalf("expected ApiKey to be 'default_key', got '%s'", cfg.ApiKey)
		}
	})

	t.Run("Multiple environment variables", func(t *testing.T) {
		os.Setenv("VAR1", "value1")
		os.Setenv("VAR2", "value2")
		defer os.Unsetenv("VAR1")
		defer os.Unsetenv("VAR2")

		type Config struct {
			First  string `env:"VAR1"`
			Second string `env:"VAR2"`
		}

		var cfg Config

		err := config.LoadConfig(&cfg)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if cfg.First != "value1" || cfg.Second != "value2" {
			t.Fatalf("expected First='value1' and Second='value2', got First='%s' Second='%s'",
				cfg.First, cfg.Second)
		}
	})

	t.Run("Empty but not required", func(t *testing.T) {
		os.Setenv("EMPTY", "")
		defer os.Unsetenv("EMPTY")

		type Config struct {
			Value string `env:"EMPTY"`
		}

		var cfg Config

		err := config.LoadConfig(&cfg)
		if err != nil {
			t.Fatalf("expected no error for empty value, got %v", err)
		}

		if cfg.Value != "" {
			t.Fatalf("expected Value to be empty, got '%s'", cfg.Value)
		}
	})
}
