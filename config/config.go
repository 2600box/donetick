package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Name             string          `mapstructure:"name" yaml:"name"`
	Telegram         TelegramConfig  `mapstructure:"telegram" yaml:"telegram"`
	Database         DatabaseConfig  `mapstructure:"database" yaml:"database"`
	Jwt              JwtConfig       `mapstructure:"jwt" yaml:"jwt"`
	Server           ServerConfig    `mapstructure:"server" yaml:"server"`
	SchedulerJobs    SchedulerConfig `mapstructure:"scheduler_jobs" yaml:"scheduler_jobs"`
	EmailConfig      EmailConfig     `mapstructure:"email" yaml:"email"`
	StripeConfig     StripeConfig    `mapstructure:"stripe" yaml:"stripe"`
	IsDoneTickDotCom bool            `mapstructure:"is_done_tick_dot_com" yaml:"is_done_tick_dot_com"`
}

type TelegramConfig struct {
	Token string `mapstructure:"token" yaml:"token"`
}

type DatabaseConfig struct {
	Type      string `mapstructure:"type" yaml:"type"`
	Host      string `mapstructure:"host" yaml:"host"`
	Port      int    `mapstructure:"port" yaml:"port"`
	User      string `mapstructure:"user" yaml:"user"`
	Password  string `mapstructure:"password" yaml:"password"`
	Name      string `mapstructure:"name" yaml:"name"`
	Migration bool   `mapstructure:"migration" yaml:"migration"`
	LogLevel  int    `mapstructure:"logger" yaml:"logger"`
}

type JwtConfig struct {
	Secret      string        `mapstructure:"secret" yaml:"secret"`
	SessionTime time.Duration `mapstructure:"session_time" yaml:"session_time"`
	MaxRefresh  time.Duration `mapstructure:"max_refresh" yaml:"max_refresh"`
}

type ServerConfig struct {
	Port             int           `mapstructure:"port" yaml:"port"`
	RatePeriod       time.Duration `mapstructure:"rate_period" yaml:"rate_period"`
	RateLimit        int           `mapstructure:"rate_limit" yaml:"rate_limit"`
	ReadTimeout      time.Duration `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout     time.Duration `mapstructure:"write_timeout" yaml:"write_timeout"`
	CorsAllowOrigins []string      `mapstructure:"cors_allow_origins" yaml:"cors_allow_origins"`
}

type SchedulerConfig struct {
	DueJob     time.Duration `mapstructure:"due_job" yaml:"due_job"`
	OverdueJob time.Duration `mapstructure:"overdue_job" yaml:"overdue_job"`
	PreDueJob  time.Duration `mapstructure:"pre_due_job" yaml:"pre_due_job"`
}

type StripeConfig struct {
	APIKey         string         `mapstructure:"api_key"`
	WhitelistedIPs []string       `mapstructure:"whitelisted_ips"`
	Prices         []StripePrices `mapstructure:"prices"`
	SuccessURL     string         `mapstructure:"success_url"`
	CancelURL      string         `mapstructure:"cancel_url"`
}

type StripePrices struct {
	PriceID string `mapstructure:"id"`
	Name    string `mapstructure:"name"`
}

type EmailConfig struct {
	Email   string `mapstructure:"email"`
	Key     string `mapstructure:"key"`
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	AppHost string `mapstructure:"appHost"`
}

func NewConfig() *Config {
	return &Config{
		Telegram: TelegramConfig{
			Token: "REMOVED",
		},
		Database: DatabaseConfig{
			Type:      "sqlite",
			Migration: true,
		},
		Jwt: JwtConfig{
			Secret:      "secret",
			SessionTime: 7 * 24 * time.Hour,
			MaxRefresh:  7 * 24 * time.Hour,
		},
	}
}

func LoadConfig() *Config {
	// read evniroment variable FORWARD_ENV and see if it's equal local then load local (config/local.yaml) using viper:
	if os.Getenv("CA_ENV") == "local" {
		viper.SetConfigName("local")
	} else if os.Getenv("CA_ENV") == "prod" {
		viper.SetConfigName("prod")
	} else {
		viper.SetConfigName("local")
	}
	// get logger and log the current environment:
	fmt.Printf("--ConfigLoad config for environment: %s ", os.Getenv("CA_ENV"))

	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	// print a useful error:
	if err != nil {
		panic(err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--ConfigLoad name : %s ", config.Name)

	return &config

	// return LocalConfig()
}