package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {

	config := Config{}

	configFile, err := getConfigFilePath()

	if err != nil {
		return config, errors.New("could not retrive home directory")
	}

	configData, err := os.ReadFile(configFile)

	if err != nil {
		return config, fmt.Errorf("could not read config: %s", configFile)
	}

	err = json.Unmarshal(configData, &config)

	if err != nil {
		return config, fmt.Errorf("could not unmarshal config")

	}

	return config, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.Current_user_name = name

	newCfg := Config{
		Db_url:            cfg.Db_url,
		Current_user_name: cfg.Current_user_name,
	}

	err := write(newCfg)

	if err != nil {
		return err
	}

	return nil

}

func write(cfg Config) error {

	cfgFilePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	dat, err := json.Marshal(cfg)

	if err != nil {
		return err
	}

	err = os.WriteFile(cfgFilePath, dat, 0600)

	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", errors.New("could not retrive home directory")
	}

	return home + "/" + configFileName, nil
}
