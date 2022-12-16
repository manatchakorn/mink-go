package config

import (
	"minkgo/internal/errs"
)

type Config interface{}

type config struct{}

func New(path string) Config {
	if y, err := openConfig(path); err != nil {
		println(err.Error())
		return nil
	} else {
		return y
	}
}

func openConfig(p string) (*config, error) {
	return nil, errs.ERROR_CONFIG_INVALID
}
