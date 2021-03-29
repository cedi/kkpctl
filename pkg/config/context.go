package config

// Context holds the current context of which cloud to use
type Context struct {
	CloudName string `yaml:"cloud"`
}

// NewContext creates a empty context for kkpctl
func NewContext() *Context {
	return &Context{
		CloudName: "imke",
	}
}
