package pkg

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"path"
)

const (
	configFile               = ".pushover"
	defaultApplicationToken  = "YOUR-APPLICATION-ID"
	defaultUserToken		 = "YOUR-USER-TOKEN"
)

var HomeDirectory string

var v = viper.New()

func init() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	HomeDirectory = home

	v.SetConfigName(configFile)
	v.SetConfigType("yaml")
	v.AddConfigPath(HomeDirectory)

	if !IsConfigFileExist() {
		_, err := os.Create(path.Join(HomeDirectory, configFile))
		if err != nil {
			panic(err)
		}
	}
}

type Config struct {
	ApplicationToken string
	UserToken string
}

//Check existence of the configuration file
func IsConfigFileExist() bool {
	if fi, err := os.Stat(path.Join(HomeDirectory, configFile)); err != nil || fi.IsDir() {
		return false
	}
	return true
}

//Read the configuration
func ReadConfigFile() (*Config, error) {
	v.SetConfigFile(path.Join(HomeDirectory, configFile))
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		ApplicationToken: v.GetString("defaultApplicationToken"),
		UserToken: v.GetString("userToken"),
	}, nil
}

//Write data of config to the configuration file
func WriteConfigFile(cfg *Config) error {
	v.Set("defaultApplicationToken", cfg.ApplicationToken)
	v.Set("userToken", cfg.UserToken)
	return v.WriteConfig()
}

//Write default config into the configuration file
func WriteDefaultConfig() error {
	return WriteConfigFile(&Config{
		ApplicationToken: defaultApplicationToken,
		UserToken: defaultUserToken,
	})
}
