package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Server struct {
		Port string `json:"host"`
	} `json:"server"`
	Redis struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"redis"`
	Options struct {
		Schema string `json:"prefix"`
		Prefix string `json:"prefix"`
	} `json:"options"`
}

func FromFile(path string) (*Config, error) {
	//opens the file specified by the path. If the file opening encounters an error,
	//it returns nil for the *Config and the error.
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//schedules the file to be closed when the surrounding function (FromFile) returns.
	//The defer statement ensures that the Close method is called
	//even if an error occurs during the function execution
	defer file.Close()

	//reads the entire content of the opened file (file) using the ReadAll function
	//from the io package. It returns a byte slice ([]byte) containing the file's content.
	//If an error occurs during reading, it returns nil for the *Config and the error.
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
