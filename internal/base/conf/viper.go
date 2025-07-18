package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	defaultConfigYamlFile string = "config"
)

var (
	loggerInitialised bool = false
)

// InitConfig initialises a viper configuration processor.
func InitConfig(appName, configYamlFile, configSubPath string) Config {
	// Defaults
	if configYamlFile == "" {
		configYamlFile = defaultConfigYamlFile
	}

	// Prepare config path
	viper.SetConfigName(configYamlFile)                    // name of config file (without extension)
	viper.SetConfigType("yaml")                            // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", appName))  // path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", appName)) // call multiple times to add many search paths
	viper.AddConfigPath(".")                               // optionally look for config in the working directory
	err := viper.ReadInConfig()                            // Find and read the config file
	if err != nil {                                        // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return viper.Sub(configSubPath)
}

// GetSubConfiguration gets the sub configuration of a configuration file.
func GetSubConfiguration(configSubPath string) Config {
	if loggerInitialised {
		panic(fmt.Errorf("fatal configuration hasn't been initialised"))
	}
	return viper.Sub(configSubPath)
}
