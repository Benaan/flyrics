package application

import (
	"github.com/spf13/viper"
)

var Config = NewManager()

func NewManager() *Manager {
	manager := &Manager{}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	setupPlatformSpecific()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return manager
}

type Manager struct {
}

func (*Manager) GetStringConfig(key string) string {
	return viper.GetString(key)
}

func (*Manager) SetConfig(name string, value interface{}) {
	viper.Set(name, value)
	if err := viper.WriteConfig(); err != nil {
		panic(err)
	}
}
