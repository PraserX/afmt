package printer

import ()

// Options are used for Server construct function.
type Options struct {
	IgnoreUnsupported bool
}

// Option specification for Printer package.
type Option func(*Options)

// OptionIgnoreUnsupported option specification.
func OptionIgnoreUnsupported(ignoreUnsupported bool) Option {
	return func(opts *Options) {
		opts.IgnoreUnsupported = ignoreUnsupported
	}
}
