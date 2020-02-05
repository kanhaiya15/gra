package cfg

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// NewConfig generates a viper configuration object which
// merges (in order from highest to lowest priority) the
// command line options, configuration file options, and
// default configuration values. This viper object becomes
// the single source of truth for the app configuration.
func NewConfig(cfgPath, cfgFile string) {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName(fmt.Sprintf("cfg.%s", cfgFile))
	viper.AddConfigPath(cfgPath)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(e.Name)
	})
}
