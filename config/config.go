package config

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type Server struct {
	Mode         string        `json:"mode"`
	HttpPort     string        `json:"http_port"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
}

type App struct {
	DefaultPageSize int    `json:"default_page_size"`
	MaxPageSize     int    `json:"max_page_size"`
	LogPath         string `json:"log_path"`
	LogFileName     string `json:"log_file_name"`
	LogFileExt      string `json:"log_file_ext"`
}

type Database struct {
	Type         string `json:"type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Project      string `json:"project"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}

func (s *Setup) ReadConfig(key string, val interface{}) error {
	if err := s.v.UnmarshalKey(key, val); err != nil {
		return errors.Wrap(err, fmt.Sprintf("The config of %s load error", key))
	}
	return nil
}
