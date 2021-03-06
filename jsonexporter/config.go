package jsonexporter

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name   string            `yaml:name`
	Path   string            `yaml:path`
	Labels map[string]string `yaml:labels`
	Type   string            `yaml:type`
	Help   string            `yaml:help`
	Values map[string]string `yaml:values`
}

func (config *Config) labelNames() []string {
	labelNames := make([]string, 0, len(config.Labels))
	for name := range config.Labels {
		labelNames = append(labelNames, name)
	}
	return labelNames
}

func loadConfig(configPath string) ([]*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config;path:<%s>,err:<%s>", configPath, err)
	}

	var configs []*Config
	if err := yaml.Unmarshal(data, &configs); err != nil {
		return nil, fmt.Errorf("failed to parse yaml;err:<%s>", err)
	}
	// Complete defaults
	for _, config := range configs {
		if config.Type == "" {
			config.Type = DefaultScrapeType
		}
		if config.Help == "" {
			config.Help = config.Name
		}
	}

	return configs, nil
}

/*
type ConfigForOP struct {
	Message string `yaml:"message"`
}


func loadConfigInit(configPath string) *ConfigForOP {
	data, err := ioutil.ReadFile(configPath)
	//check(err)

	var configs []*Config
	//check(err)

	endpointURL := configs[0].Labels["endpoint"]

	var conf ConfigForOP
	conf.Message = endpointURL

	return &conf
}*/

func loadConfigInit(configPath string) (string, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to load config;path:<%s>,err:<%s>", configPath, err)
	}

	var configs []*Config
	if err := yaml.Unmarshal(data, &configs); err != nil {
		return "", fmt.Errorf("failed to parse yaml;err:<%s>", err)
	}

	endpointURL := configs[0].Labels["endpoint"]

	return endpointURL, nil
}
