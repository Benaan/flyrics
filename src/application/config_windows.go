package application

import (
	"github.com/benaan/flyrics/src/config"
	"github.com/spf13/viper"
)

func setupPlatformSpecific() {
	viper.AddConfigPath("$HOMEDRIVE$HOMEPATH\\AppData\\Roaming\\Flyrics")
	viper.SetDefault(config.GpmdpPath, "$HOMEDRIVE$HOMEPATH\\AppData\\Roaming\\Google Play Music Desktop Player\\json_store\\playback.json")
	viper.SetDefault(config.LyricDirectory, "$HOMEDRIVE$HOMEPATH\\Music\\Lyrics")
}
