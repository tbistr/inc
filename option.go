package inc

// option for Engine.
type option struct {
	ignoreCase bool
}

// Option is a function that configures the option.
type Option func(*option)

// IgnoreCase makes the Engine ignore case.
func IgnoreCase() Option {
	return func(o *option) {
		o.ignoreCase = true
	}
}
