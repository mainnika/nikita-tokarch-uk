package config

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Backend contains backend connection-specific configuration
type Backend struct {
	Addr    string `mapstructure:"addr"`
	Secured bool   `mapstructure:"secured"`
}

// Content contains content-specific configuration
type Content struct {
	Backend      Backend `mapstructure:"backend"`
	Key          string  `mapstructure:"key"`
	Pinned       string  `mapstructure:"pinned"`
	PostsPerPage int     `mapstructure:"postsPerPage"`
}

// Config contains application configuration
type Config struct {
	Verbose string  `mapstructure:"verbose"`
	Base    string  `mapstructure:"base"`
	Addr    string  `mapstructure:"addr"`
	Unix    string  `mapstructure:"unix"`
	Content Content `mapstructure:"content"`
}

// initialize default values on app-start
func init() {

	pflag.BoolP("version", "V", false, "get application version")
	pflag.BoolP("verbose", "v", false, "enable verbose logging")
	pflag.StringP("config", "c", string("frontend.yaml"), "config file path")

	pflag.String("addr", "127.0.0.1:8000", "tcp addr to listen")
	pflag.String("unix", "", "unix socket path to listen")
	pflag.String("base", "", "http URI prefix")

	pflag.String("content.backend.addr", "demo.ghost.io:443", "ghost backend addr")
	pflag.Bool("content.backend.secured", true, "is ghost backend secured")
	pflag.String("content.backend", "demo.ghost.io:443", "ghost backend addr")
	pflag.String("content.key", "22444f78447824223cefc48062", "ghost content api key")
	pflag.String("content.pinned", "contact", "pinned page slug")
	pflag.Int("content.postsPerPage", 5, "amount of posts per page")

	pflag.Parse()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.BindPFlags(pflag.CommandLine)

	if viper.GetBool("verbose") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Verbose mode")
	}
}
