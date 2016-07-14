package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/AndreaGhizzoni/go-status/constant"
	"github.com/AndreaGhizzoni/go-status/util"
)

// TODO add doc
type Config struct {
	Watchedpaths []string `json:"wp"`
}

var (
	defCfg string = "{ \"wp\" : [] }"
)

// TODO add doc
func Parse() (*Config, error) {
	c := &Config{}

	if e, _ := util.ExistsPath(constant.ConfigPath); !e {
		err := ioutil.WriteFile(constant.ConfigPath,
			[]byte(defCfg), 0644)
		if err != nil {
			return nil, err
		}
	} else {
		file, err := ioutil.ReadFile(constant.ConfigPath)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(file, c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
