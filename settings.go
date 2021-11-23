package main

import (
	"os"
	"path"

	"github.com/cacilhas/sudogo/sudoku"
	"github.com/spf13/viper"
)

func readSettings() {
	viper.SetConfigName("sudogo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path.Join(os.Getenv("HOME"), ".config"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}
	defaultSettings()
	viper.Set("homepage", "https://cacilhas.itch.io/nonogram")
	viper.Set("version", "nightly")
}

func defaultSettings() {
	if !viper.IsSet("hardship") {
		viper.Set("hardship", int(sudoku.MEDIUM))
	}
	if viper.GetInt("width") == 0 {
		viper.Set("width", 600)
	}
	if viper.GetInt("height") == 0 {
		viper.Set("height", 600)
	}
}

func saveSettings() {
	if err := viper.WriteConfig(); err != nil {
		if err = viper.SafeWriteConfig(); err != nil {
			panic(err)
		}
	}
}
