package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ProjectName    string `default:"eh-digital-shift"`
	ProjectVersion string `envconfig:"VERSION" default:"0.0.1"`
	Prefix         string `envconfig:"PREFIX" default:"/digital-shift"`
	Port           string `envconfig:"PORT" default:"8000"`
	Env            string `envconfig:"ENV" default:"develop"`
	Mongo          struct {
		Uri        string `envconfig:"MONGO_URI" default:"mongodb://root:example@mongo:27017/"`
		Database   string `envconfig:"MONGO_DATABASE" default:"digital_shift"`
		Collection string `envconfig:"MONGO_COLLECTION" default:"booking"`
	}
}

func NewConfig() *Config {
	var conf Config
	if err := envconfig.Process("", &conf); err != nil {
		panic(err.Error())
	}

	return &conf
}
