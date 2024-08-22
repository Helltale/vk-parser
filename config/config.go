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
	DownloadLink    string `yaml:"download-link"`
}

func (c *Config) GetResultDirectory() string {
	return c.ResultDirectory
}

func (c *Config) GetClassNameImg() string {
	return c.ClassNameImg
}

func (c *Config) GetDownloadLink() string {
	return c.DownloadLink
}

func (c *Config) SetConfig(config []byte) error {
	return yaml.Unmarshal(config, c)
}

type Api struct {
	ApiToken   string `yaml:"api-token"`
	ApiKey     string `yaml:"api-key"`
	ApiVersion string `yaml:"api-version"`
}

func (a *Api) GetApiToken() string {
	return a.ApiToken
}

func (a *Api) GetApiKey() string {
	return a.ApiKey
}

func (a *Api) GetApiVersion() string {
	return a.ApiVersion
}

func (a *Api) SetConfig(config []byte) error {
	return yaml.Unmarshal(config, a)
}

type Configurable interface {
	SetConfig(config []byte) error
}

func getYAMLPath(file string) (string, error) {
	projectDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error: can not get project directory: %v", err)
	}
	projectDir = filepath.Dir(projectDir)

	filePath := filepath.Join(projectDir, "config", file)

	return filePath, nil
}

// for example:
// config - api
// nameFile - api.yaml
func getConfig(c Configurable, nameFile string) error {
	filePath, err := getYAMLPath(nameFile)
	if err != nil {
		return errors.New("error: cannot get path file " + nameFile)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return errors.New("error: cannot get config file " + nameFile)
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		return errors.New("error: cannot read config file " + nameFile)
	}

	err = c.SetConfig(read)
	if err != nil {
		return errors.New("error: cannot unmarshal config file " + nameFile)
	}

	return nil
}

func GetCassNameImg() (string, error) {
	var cfg Config
	err := getConfig(&cfg, "config.yaml")
	if err != nil {
		return "", err
	}
	return cfg.GetClassNameImg(), nil
}

func GetResultDirectory() (string, error) {
	var cfg Config
	err := getConfig(&cfg, "config.yaml")
	if err != nil {
		return "", err
	}
	return cfg.GetResultDirectory(), nil
}

func GetDownloadLink() (string, error) {
	var cfg Config
	err := getConfig(&cfg, "config.yaml")
	if err != nil {
		return "", err
	}
	return cfg.GetDownloadLink(), nil
}

func GetApiToken() (string, error) {
	var api Api
	err := getConfig(&api, "api.yaml")
	if err != nil {
		return "", err
	}
	return api.GetApiToken(), nil
}

func GetApiKey() (string, error) {
	var api Api
	err := getConfig(&api, "api.yaml")
	if err != nil {
		return "", err
	}
	return api.GetApiKey(), nil
}

func GetApiVersion() (string, error) {
	var api Api
	err := getConfig(&api, "api.yaml")
	if err != nil {
		return "", err
	}
	return api.GetApiVersion(), nil
}

// Функция для чтения YAML файла в map[string]interface{}
func readYAMLFile(filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error: cannot get config file %s: %v", filepath.Base(filePath), err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error: cannot read config file %s: %v", filepath.Base(filePath), err)
	}

	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error: cannot unmarshal config file %s: %v", filepath.Base(filePath), err)
	}

	return config, nil
}

// Функция для записи map[string]interface{} в YAML файл
func writeYAMLFile(filePath string, config map[string]interface{}) error {
	updatedData, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("error: cannot marshal config file %s: %v", filepath.Base(filePath), err)
	}

	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("error: cannot write config file %s: %v", filepath.Base(filePath), err)
	}

	return nil
}

// Функция для обновления значения поля в YAML файле по ключу
func UpdateYAMLField(filePath string, key string, newValue interface{}) error {
	filePath1, err := getYAMLPath(filePath)
	if err != nil {
		return err
	}

	config, err := readYAMLFile(filePath1)
	if err != nil {
		return err
	}

	config[key] = newValue

	err = writeYAMLFile(filePath1, config)
	if err != nil {
		return err
	}

	return nil
}

// Функция для вывода содержимого YAML файла
func PrintYAMLFile(filePath string) error {
	filePath1, err := getYAMLPath(filePath)
	if err != nil {
		return err
	}

	config, err := readYAMLFile(filePath1)
	if err != nil {
		return err
	}

	fmt.Println("config content:")
	for key, value := range config {
		fmt.Printf("%v: %v\n", key, value)
	}

	return nil
}
