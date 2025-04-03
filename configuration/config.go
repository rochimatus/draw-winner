package configuration

import (
	"fmt"
	"github.com/goto/salt/config"
	"github.com/rochimatus/draw-winner/logger"
	"time"
)

type Configuration struct {
	HTTP HTTP `mapstructure:"http" json:"http" yaml:"http"`
}

type HTTP struct {
	Port                     string        `mapstructure:"port" default:"8080" json:"port" yaml:"port"`
	GraceFulShutDownDuration time.Duration `mapstructure:"graceful_shutdown_duration" default:"20s" json:"graceful_shutdown_duration" yaml:"graceful_shutdown_duration"`
	ReadHeaderTimeout        time.Duration `mapstructure:"read_header_timeout" default:"2s" json:"read_header_timeout" yaml:"read_header_timeout"`
}

func LoadConfig(path string) (*Configuration, error) {
	var conf Configuration
	configLoader := config.NewLoader(
		config.WithPath(path),
	)

	if err := configLoader.Load(&conf); err != nil {
		return nil, fmt.Errorf("unable to load config: %w", err)
	}

	logger.Info("Loaded Env configs are:", "config", conf)

	return &conf, nil
}
