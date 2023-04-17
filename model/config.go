package model

import (
	"encoding/json"
	"github.com/tcnksm/go-input"
	"log"
	"os"
)

var Conf Config

var ui = &input.UI{
	Writer: log.Writer(),
	Reader: os.Stdin,
}

func LoadConfig() error {
	// Try to load config from file
	file, err := os.ReadFile("config.json")
	if err == nil {
		err = json.Unmarshal(file, &Conf)
		return err
	}

	// If file doesn't exist, ask user to input config values
	Conf.ApiKey, err = ui.Ask("API key", &input.Options{
		HideOrder: true,
		Required:  true,
		Loop:      true,
	})
	if err != nil {
		return err
	}

	Conf.OaCookie, err = ui.Ask("OA cookie", &input.Options{
		HideOrder: true,
		Required:  true,
		Loop:      true,
	})
	if err != nil {
		return err
	}
	// Save config to file
	file, err = json.MarshalIndent(Conf, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("config.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}
