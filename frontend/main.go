package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	logrus.Infof("Version: %s", Version)
}
