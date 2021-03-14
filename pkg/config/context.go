package config

// Context holds the current context of which cloud to use
type Context struct {
	CloudName string `yaml:"cloud"`
	Bearer    string `yaml:"bearer"`
}

// NewContext creates a empty context for kkpctl
func NewContext() Context {
	return Context{
		CloudName: "",
		Bearer:    "",
	}
}
