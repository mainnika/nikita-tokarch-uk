package config

// Content contains content-specific configuration
type Content struct {
	Backend      string `mapstructure:"backend"`
	Key          string `mapstructure:"key"`
	Pinned       string `mapstructure:"pinned"`
	PostsPerPage int    `mapstructure:"postsPerPage"`
}

// Config contains application configuration
type Config struct {
	Verbose string  `mapstructure:"verbose"`
	Base    string  `mapstructure:"base"`
	Addr    string  `mapstructure:"addr"`
	Unix    string  `mapstructure:"unix"`
	Content Content `mapstructure:"content"`
}
