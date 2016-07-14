package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/AndreaGhizzoni/go-status/constant"
)

// TODO add doc
type Config struct {
	Watchedpaths []string `json:"wp"`
}

// TODO add doc
func New() (*Config, error) {
	file, err := ioutil.ReadFile(constant.ConfigPath)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = json.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
