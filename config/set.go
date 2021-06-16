package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Setup config basic setup
type Setup struct {
	v *viper.Viper
}

func NewSetup() (*Setup, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("config/")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "setup new error")
	}
	return &Setup{v}, nil
}
