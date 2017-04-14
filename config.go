package hkbridge

import (
	"encoding/json"
	"io/ioutil"
)

type AccessoryConfig struct {
	Type   string                 `json:"type"`
	Option map[string]interface{} `json:"option"`
}

type Config struct {
	Name         string            `json:"name"`
	Manufacturer string            `json:"manufacturer"`
	SerialNumber string            `json:"serial"`
	Model        string            `json:"model"`
	Pin          string            `json:"pin"`
	Port         string            `json:"port"`
	Interface    string            `json:"interface"`
	Debug        bool              `json:"debug"`
	Accessories  []AccessoryConfig `json:"accessories"`
}

func ReadConfig(file string) (Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(data, &config)
	return config, err
}
