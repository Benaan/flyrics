package application

import (
	"github.com/benaan/flyrics/src/config"
	"github.com/spf13/viper"
)

func setupPlatformSpecific() {
	viper.AddConfigPath("$HOME/.config/flyrics")
	viper.SetDefault(config.LyricDirectory, "$HOME/Music/Lyrics")
	viper.SetDefault(config.GpmdpPath, "$HOME/.config/Google Play Music Desktop Player/json_store/playback.json")
}
