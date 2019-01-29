// Package afmt (Advanced formatter or Advanced fmt) implement some extensions
// for fmt package. The main feature is print to tree structure.
package afmt

// TreePrinterOptions are used for Server construct function.
type TreePrinterOptions struct {
	// Ignore unsupported types, otherwise return error
	IgnoreUnsupported bool
	// Padding by spaces
	Padding uint
	// Pretty name for printer
	PrettyNames map[string]string
	// Ignore specified parts of structure (do not print these ones)
	IgnoreNames []string
}

// TreePrinterOption specification for Printer package.
type TreePrinterOption func(*TreePrinterOptions)

// TreePrinterOptionIgnoreUnsupported option specification.
func TreePrinterOptionIgnoreUnsupported(ignoreUnsupported bool) TreePrinterOption {
	return func(opts *TreePrinterOptions) {
		opts.IgnoreUnsupported = ignoreUnsupported
	}
}

// TreePrinterOptionPadding option specification.
func TreePrinterOptionPadding(padding uint) TreePrinterOption {
	return func(opts *TreePrinterOptions) {
		opts.Padding = padding
	}
}

// TreePrinterOptionPrettyNames option specification.
func TreePrinterOptionPrettyNames(names map[string]string) TreePrinterOption {
	return func(opts *TreePrinterOptions) {
		opts.PrettyNames = names
	}
}

// TreePrinterOptionIgnoreNames option specification.
func TreePrinterOptionIgnoreNames(names []string) TreePrinterOption {
	return func(opts *TreePrinterOptions) {
		opts.IgnoreNames = names
	}
}
