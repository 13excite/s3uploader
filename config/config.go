package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// DefaultConfigPath is path to yaml config with key, backets and etc, usage if have not -c --config key when run script
const DefaultConfigPath = "/etc/s3config.yaml"

// Config structure described yaml structure from config.yaml
type Config struct {
	Listen      string `yaml:"listen"`
	URL         string `yaml:"url"`
	BucketName  string `yaml:"bucket"`
	Expire      int    `yaml:"expire"`
	Credentinal struct {
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
	}
}

//GetConf reading and parsing configuration file
func (conf *Config) GetConf(filePath string) {
	if filePath == "" {
		filePath = DefaultConfigPath
	}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("yamlFile read fail with err: #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println("Unmarshal error:")
		log.Fatal(err)
		os.Exit(1)
	}
}
