package configuration

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

// DefaultConfigurationFilename defines the default configuration file name.
const DefaultConfigurationFilename = ".config/go-chess-xfcc/go-chess-xfcc.xml"

// Configuration is simple configuration for the username and the password.
type Configuration struct {
	User     string   `xml:"user"`
	Password string   `xml:"password"`
	XMLName  xml.Name `xml:"configuration"`
}

// GetConfiguration returns the configuration from the default location.
func GetConfiguration() (Configuration, error) {
	config := Configuration{}
	configurationFile, err := getConfigurationFilename()
	if err != nil {
		return config, err
	}

	if err := config.Load(configurationFile); err != nil {
		return config, err
	}

	return config, nil
}

// Load loads the configuration from an XML file
func (c *Configuration) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(bytes, c)
	if err != nil {
		return err
	}

	return nil
}

func getHomeDirectory() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

func getConfigurationFilename() (string, error) {
	home, err := getHomeDirectory()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", home, DefaultConfigurationFilename), nil
}
