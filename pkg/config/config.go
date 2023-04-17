package config

import "github.com/spf13/viper"

type Config struct {
	TelegramToken     string `mapstructure: "token"`
	UnsplashAccessKey string `mapstructure: "access_key"`
}

func Init() (*Config, error) {
	if err := loadConfig(); err != nil {
		return nil, err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
func loadConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	
	return viper.ReadInConfig()
}
func fromEnv(cfg *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("token")

	if err := viper.BindEnv("access_key"); err != nil {
		return err
	}
	cfg.UnsplashAccessKey = viper.GetString("access_key")

	return nil
}
