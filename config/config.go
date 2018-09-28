package config

import "github.com/caarlos0/env"

// Config struct to hold all configurable parameters
type Config struct {
	ServiceName     string `env:"SERVICE_NAME" envDefault:"chupapintos_translate_bot"`
	TranslateApiKey string `env:"TRANSLATE_API_KEY" envDefault:""`
	TelegramBotKey  string `env:"TELEGRAM_BOT_KEY" envDefault:""`
	TranslateApiURL string `env:"TELEGRAM_BOT_KEY" envDefault:"https://translate.yandex.net/api/v1.5/tr.json/translate"`
	LangugeSrc      string `env:"LANGUAGE_SRC" envDefault:"pt"`
	LanguageDst     string `env:"LANGUAGE_DST" envDefault:"en"`
}

func New() (Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return cfg, err
}
