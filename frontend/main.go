package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(viper.GetString("config"))
	if err := viper.ReadInConfig(); err != nil {
		logrus.Debugf("Skip invalid config file %s, %v", viper.ConfigFileUsed(), err)
	}
}
