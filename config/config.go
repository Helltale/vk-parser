package config

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ResultDirectory string `yaml:"result-directory"`
	ClassNameImg    string `yaml:"class-name-img"`
	DownloadLink    string `yaml:"download-link"`
}

type Api struct {
	ApiToken   string `yaml:"api-token"`
	ApiKey     string `yaml:"api-key"`
	ApiVersion string `yaml:"api-version"`
}

func getConf() (*Config, error) {
	projectDir, err := os.Getwd()
	if err != nil {
		return nil, errors.New("error: can not get project directory")
	}
	// Получаем директорию, в которой находится исполняемый файл
	projectDir = filepath.Dir(projectDir)

	file, err := os.Open(filepath.Join(projectDir, "config", "config.yaml"))
	if err != nil {
		return nil, errors.New("error: cannot get config file config.yaml")
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("error: cannot read config file config.yaml")
	}

	c := Config{}

	err = yaml.Unmarshal(read, &c)
	if err != nil {
		return nil, errors.New("error: cannot unmarshal config file config.yaml")
	}

	return &c, nil
}

func getApi() (*Api, error) {
	projectDir, err := os.Getwd()
	if err != nil {
		return nil, errors.New("error: can not get project directory")
	}
	// Получаем директорию, в которой находится исполняемый файл
	projectDir = filepath.Dir(projectDir)

	file, err := os.Open(filepath.Join(projectDir, "config", "config.yaml"))
	if err != nil {
		return nil, errors.New("error: cannot get config file config.yaml")
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("error: cannot read config file config.yaml")
	}

	api := Api{}

	err = yaml.Unmarshal(read, &api)
	if err != nil {
		return nil, errors.New("error: cannot unmarshal config file config.yaml")
	}

	return &api, nil
}

func GetCassNameImg() (string, error) {
	conf, err := getConf()
	if err != nil {
		return "", err
	}
	return conf.ClassNameImg, nil
}

func GetResultDirectory() (string, error) {
	conf, err := getConf()
	if err != nil {
		return "", err
	}
	return conf.ResultDirectory, nil
}

func GetDownloadLink() (string, error) {
	conf, err := getConf()
	if err != nil {
		return "", err
	}
	return conf.DownloadLink, nil
}

func GetApiToken() (string, error) {
	api, err := getApi()
	if err != nil {
		return "", err
	}
	return api.ApiToken, nil
}

func GetApiKey() (string, error) {
	api, err := getApi()
	if err != nil {
		return "", err
	}
	return api.ApiKey, nil
}

func GetApiVersion() (string, error) {
	api, err := getApi()
	if err != nil {
		return "", err
	}
	return api.ApiVersion, nil
}
