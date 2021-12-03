package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"

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

	var netw, addr string
	switch {
	case config.Addr != "":
		netw = "tcp"
		addr = config.Addr
	case config.Unix != "":
		netw = "unix"
		addr = config.Unix
	default:
		panic("no address given")
	}
	netListener, err := net.Listen(netw, addr)
	if err != nil {
		logrus.Fatal(err)
	}

	httpServer := fasthttp.Server{
		Logger:  logrus.StandardLogger(),
		GetOnly: true,
	}

	logrus.Infof("Version: %s", Version)
	logrus.Debugf("Addr: %s+%s", netw, addr)
	logrus.Debugf("Conf: %#v", config)

	err = httpServer.Serve(netListener)
	if err != nil {
		logrus.Fatal(err)
	}
}
