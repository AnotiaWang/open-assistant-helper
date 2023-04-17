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

func askForOpenAIKey() (string, error) {
	return ui.Ask("Please input your OpenAI API key", &input.Options{
		HideOrder: true,
		Required:  true,
		Loop:      true,
	})
}

func askForOpenAssistantCookie() (string, error) {
	return ui.Ask("Please input your Open Assistant cookie", &input.Options{
		HideOrder: true,
		Required:  true,
		Loop:      true,
	})
}

func askForLanguage() (string, error) {
	return ui.Ask("Please input your preferred language for the tasks", &input.Options{
		Default:   "zh",
		HideOrder: true,
		Loop:      true,
	})
}

func LoadConfig() error {
	// Try to load config from file
	file, err := os.ReadFile("config.json")
	if err == nil {
		err = json.Unmarshal(file, &Conf)
		if err != nil {
			return err
		}
		if Conf.ApiKey == "" {
			Conf.ApiKey, err = askForOpenAIKey()
			if err != nil {
				return err
			}
		}
		if Conf.OaCookie == "" {
			Conf.OaCookie, err = askForOpenAssistantCookie()
			if err != nil {
				return err
			}
		}
		if Conf.Language == "" {
			Conf.Language, err = askForLanguage()
			if err != nil {
				return err
			}
		}
		file, _ = json.MarshalIndent(Conf, "", "    ")
		_ = os.WriteFile("config.json", file, 0644)
		return nil
	}

	// If file doesn't exist, ask user to input config values
	Conf.ApiKey, err = askForOpenAIKey()
	if err != nil {
		return err
	}

	Conf.OaCookie, err = askForOpenAssistantCookie()
	if err != nil {
		return err
	}
	Conf.Language, err = askForLanguage()
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

func UpdateCookie(cookie string) error {
	Conf.OaCookie = cookie
	file, err := json.MarshalIndent(Conf, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("config.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
