package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const defaultConfigFile = "app.yaml"

type AppConfig struct {
	DBTable string `yaml:"dbTable"`
}

func GetAppConfig() (*AppConfig, error) {
	cfg, err := FromYaml(defaultConfigFile, "dev")
	if err != nil {
		return nil, errors.New("failed to get app config")
	}
	return cfg, nil
}

func FromYaml(fileName, evn string) (*AppConfig, error) {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("failed to read yaml file")
	}

	appConfigs := make(map[string]AppConfig)
	if err := yaml.Unmarshal(fileContent, &appConfigs); err != nil {
		return nil, errors.New("failed to read yaml file")
	}

	appCfg, _ := appConfigs["dev"]

	return &appCfg, nil
}
