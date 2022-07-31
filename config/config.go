package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Telegram struct {
		Token string `yaml:"token"`
	} `yaml:"telegram"`
	YouTube struct {
		YoutubeUrl string `yaml:"youtube_url"`
		ApiToken   string `yaml:"api_token"`
	} `yaml:"youtube"`
	ChuckNorris struct {
		ChuckUrl string `yaml:"chuck_url"`
	} `yaml:"chucknorris"`
	Weather struct {
		WeatherApi string `yaml:"weather_api"`
		WeatherKey string `yaml:"weather_key"`
	} `yaml:"weather"`
}

var cfg Config

func GetConfig() *Config {
	cfg := &Config{}

	err := cleanenv.ReadConfig("config.yml", cfg)
	if err != nil {
		return nil
	}

	return cfg
}
