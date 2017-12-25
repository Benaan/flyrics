package application

import (
	"github.com/benaan/flyrics/src/config"
	"github.com/spf13/viper"
)

func setupPlatformSpecific() {
	viper.AddConfigPath("~\\AppData\\Roaming\\Flyrics")
	viper.SetDefault(config.LyricDirectory, "~\\AppData\\Roaming\\Google Play Music Desktop Player\\json_store\\playback.json")
	viper.SetDefault(config.GpmdpPath, "~\\My Music\\Lyrics")
}
