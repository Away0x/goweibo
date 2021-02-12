package tpl

// Config tpl config
type Config struct {
	GetRoutePath func(name string, params ...interface{}) string
	GeneratePublicPath func(path string) string
}

var config *Config

// SetupTpl setup tpl
func SetupTpl(c *Config) {
	config = c
}
