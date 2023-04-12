package inc

import "strings"

// config for Engine.
type config struct {
	index Index
}

// defaultConfig
var defaultConfig = &config{
	index: strings.IndexRune,
}

// option is a function that configures the option.
type option func(*config)

// IgnoreCase makes the Engine ignore case.
func IgnoreCase() option {
	return UseCustomIndex(indexIgnoreCase)
}

// UseCustomIndex makes the Engine use a custom search function.
func UseCustomIndex(index Index) option {
	return func(c *config) {
		c.index = index
	}
}
