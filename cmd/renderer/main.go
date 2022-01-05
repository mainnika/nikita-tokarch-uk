package main

import (
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"

	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/config"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/ghost/v4api/httpclient"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/routes"
	"code.tokarch.uk/mainnika/nikita-tokarch-uk/pkg/templates"
)

var Version = "nightly"

const frontendServerIdentity = "Microsoft-IIS/6.0"

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

	templateFuncs := &templates.Funcs{Version: Version, Site: config.Site}
	err = templates.Load(templateFuncs)
	if err != nil {
		logrus.Fatal(err)
	}

	ghostClient := &httpclient.HTTPClient{
		Addr:         config.Content.Backend.Addr,
		Secured:      config.Content.Backend.Secured,
		Headers:      config.Content.Backend.Headers,
		ContentKey:   config.Content.Key,
		QueryTimeout: time.Second,
	}

	apiRoutes := &routes.Routes{
		GhostClient:   ghostClient,
		ContentConfig: config.Content,
		Base:          config.Base,
	}

	httpServer := fasthttp.Server{
		Logger:  logrus.StandardLogger(),
		Handler: apiRoutes.Handler,
		Name:    frontendServerIdentity,
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
