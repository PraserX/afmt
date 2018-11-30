package afmt

import ()

// PrinterOptions are used for Server construct function.
type PrinterOptions struct {
	// Ignore unsupported types, otherwise return error
	IgnoreUnsupported bool
	// Padding by spaces
	Padding uint
	// Pretty name for printer
	PrettyNames map[string]string
}

// PrinterOption specification for Printer package.
type PrinterOption func(*PrinterOptions)

// PrinterOptionIgnoreUnsupported option specification.
func PrinterOptionIgnoreUnsupported(ignoreUnsupported bool) PrinterOption {
	return func(opts *PrinterOptions) {
		opts.IgnoreUnsupported = ignoreUnsupported
	}
}

// PrinterOptionPadding option specification.
func PrinterOptionPadding(padding uint) PrinterOption {
	return func(opts *PrinterOptions) {
		opts.Padding = padding
	}
}

// PrinterOptionPrettyNames option specification.
func PrinterOptionPrettyNames(names map[string]string) PrinterOption {
	return func(opts *PrinterOptions) {
		opts.PrettyNames = names
	}
}
