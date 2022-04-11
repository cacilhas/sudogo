package ui

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
	"github.com/sqweek/dialog"
)

var filters [][]string = [][]string{
	{"Sudoku files", "sd"},
	{"Plain text", "txt"},
	{"Any"},
}

func getDialog(title string, filter bool) *dialog.FileBuilder {
	resp := dialog.File().SetStartDir(viper.GetString("save_dir")).Title(title)
	if filter {
		for _, filt := range filters {
			resp.Filter(filt[0], filt[1:]...)
		}
	}
	return resp
}

func showError(err error) {
	showMessage(err.Error(), "error")
}

func showInfo(msg string, args ...interface{}) {
	showMessage(fmt.Sprintf(msg, args...), "info")
}

func showMessage(msg string, tpe string) {
	builder := dialog.Message(msg).Title("Sudoku")
	switch tpe {
	case "error":
		builder.Error()
	default:
		builder.Info()
	}
}

func fixFilename(filename string) string {
	if filename == "" {
		return ""
	}
	if !path.IsAbs(filename) {
		filename = path.Join(viper.GetString("save_dir"), filename)
	}
	if path.Ext(filename) == "" {
		filename += ".sd"
	}
	return filename
}

func openFile() (*os.File, error) {
	builder := getDialog("Load Board", true)
	if filename, err := getFilename(builder.Load); err == nil {
		return os.Open(filename)
	} else {
		return nil, err
	}
}

func saveFile() (*os.File, error) {
	builder := getDialog("Save Board", false)
	if filename, err := getFilename(builder.Save); err == nil {
		return os.Create(filename)
	} else {
		return nil, err
	}
}

func getFilename(read func() (string, error)) (string, error) {
	if filename, err := read(); err == nil {
		filename = fixFilename(filename)
		if filename == "" {
			return "", fmt.Errorf("no file selected")
		}
		viper.Set("save_dir", path.Dir(filename))
		return filename, nil
	} else {
		return "", err
	}
}
