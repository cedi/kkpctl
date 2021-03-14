package config

// Context holds the current context of which cloud to use
type Context struct {
	URL    string `yaml:"url"`
	Bearer string `yaml:"bearer"`
}

func NewContext() Context {
	return Context{
		URL:    "",
		Bearer: "",
	}
}
