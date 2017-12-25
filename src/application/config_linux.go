package application

import (
	"github.com/benaan/flyrics/src/config"
	"github.com/spf13/viper"
)

func setupPlatformSpecific() {
	viper.AddConfigPath("~/.config/flyrics")
	viper.SetDefault(config.LyricDirectory, "~/Music/Lyrics")
	viper.SetDefault(config.GpmdpPath, "~/.config/Google Play Music Desktop Player/json_store/playback.json")
}
