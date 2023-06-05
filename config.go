package main

import (
	"fmt"
)

/*
 * This is a global, Singleton Config object to hold the defaults and dynamically
 * configured values for how we should run all in one place.
 */
type Config struct {
	data map[string]string
}

var instance *Config

func GetGlobalConfig() *Config {
	if instance == nil {
		instance = newConfig()
	}

	return instance
}

func newConfig() *Config {
	config := new(Config)

	// Default values
	config.data = map[string]string{
		"listenPort":       "5154",
		"obfuscationKey":   "ed91b052346e31d5567820edde46a641",
		"rendezvousServer": "henfred.hopto.org",
	}

	return config
}

func (c *Config) Set(key string, value string) {
	c.data[key] = value
}

func (c *Config) Get(key string) string {
	return c.data[key]
}

func (c *Config) String() string {
	var str string
	for key, value := range c.data {
		str += fmt.Sprintf("\"%s\" -> \"%s\"\n", key, value)
	}

	return str
}
