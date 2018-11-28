package printer

import ()

// Options are used for Server construct function.
type Options struct {
	// Ignore unsupported types, otherwise return error
	IgnoreUnsupported bool
	// Padding by spaces
	Padding uint
}

// Option specification for Printer package.
type Option func(*Options)

// OptionIgnoreUnsupported option specification.
func OptionIgnoreUnsupported(ignoreUnsupported bool) Option {
	return func(opts *Options) {
		opts.IgnoreUnsupported = ignoreUnsupported
	}
}

// OptionPadding option specification.
func OptionPadding(padding uint) Option {
	return func(opts *Options) {
		opts.Padding = padding
	}
}
