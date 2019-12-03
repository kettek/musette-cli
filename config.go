package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Server string
}

func (c *Config) LoadFromFile(filePath string) (err error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, c)
	return err
}
