package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/frontend/config"
)

var Version = "nightly"

func main() {

	if viper.GetBool("version") {
		fmt.Print(Version)
		return
	}

	viper.SetConfigFile(viper.GetString("config"))
	if err := viper.ReadInConfig(); err != nil {
		logrus.Debugf("Skip invalid config file %s, %v", viper.ConfigFileUsed(), err)
	}

	config := &config.Config{}
	if err := viper.Unmarshal(config); err != nil {
		logrus.Warnf("Cannot unmarshal config, %v", err)
	}

	logrus.Infof("Version: %s", Version)
	logrus.Debugf("Conf: %#v", config)
}
