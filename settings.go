package main

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

var homedir string

func init() {
	var err error
	if homedir, err = os.UserHomeDir(); err != nil {
		homedir = os.Getenv("HOME")
	}

	viper.SetConfigName("sudogo")
	viper.SetConfigType("yaml")
	config := os.Getenv("XDG_CONFIG_HOME")
	if config == "" {
		config = path.Join(homedir, ".config")
	}
	viper.AddConfigPath(config)
}

func readSettings() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}
	defaultSettings()
	viper.Set("homepage", "https://cacilhas.itch.io/sudogo")
	viper.Set("version", "1.2")
}

func saveSettings() {
	fmt.Println("saving settings…")
	if err := viper.WriteConfig(); err != nil {
		if err = viper.SafeWriteConfig(); err != nil {
			panic(err)
		}
	}
}

func defaultSettings() {
	if viper.GetInt("width") == 0 {
		viper.Set("width", 1280) // TODO: get resolution
	}
	if viper.GetInt("height") == 0 {
		viper.Set("height", 720) // TODO: get resolution
	}
	if !viper.IsSet("fullscreen") {
		viper.Set("fullscreen", false)
	}
	if !viper.IsSet("save_dir") {
		viper.Set("save_dir", homedir)
	}
}
