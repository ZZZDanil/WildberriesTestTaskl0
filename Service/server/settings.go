package service

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type DataBaseOptions struct {
	User     string
	Password string
	DBName   string
	SSLMode  string
}
type MessageBrokerOptions struct {
	ClusterID string
	ClientID  string
	Url       string
}
type ServiceOptions struct {
	Host string
	Port string
}
type SettingsOptions struct {
	DataBase      DataBaseOptions
	MessageBroker MessageBrokerOptions
	Service       ServiceOptions
}

func (s *SettingsOptions) LoadSettings(path string) {
	pathToSettings, _ := filepath.Abs(path)
	jsonFile, _ := os.ReadFile(pathToSettings)
	json.Unmarshal(jsonFile, &s)
}
