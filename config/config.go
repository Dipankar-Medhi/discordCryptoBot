// THis file will bring the vars from config.json to the application

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

// we will read the config.json file and get them in json
type configStruct struct {
	Token     string `json: "Token"`
	BotPrefix string `josn: "BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	// unmarshal the file and bring the variables in config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
