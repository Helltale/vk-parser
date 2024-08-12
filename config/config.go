package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ResultDirectory string `yaml:"result-directory"`
	ClassNameImg    string `yaml:"class-name-img"`
}

// func getConfig() {

// }

func getConf() (*Config, error) {
	projectDir, err := os.Getwd()
	if err != nil {
		return nil, errors.New("error: can not get project directory\n")
	}
	// Получаем директорию, в которой находится исполняемый файл
	projectDir = filepath.Dir(projectDir)
	fmt.Println(projectDir)

	file, err := os.Open(filepath.Join(projectDir, "config", "config.yaml"))
	if err != nil {
		return nil, errors.New("error: cannot get config file config.yaml.")
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
