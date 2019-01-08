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

// PrinterOptionIgnoreUnsupported option specification.
func PrinterOptionIgnoreUnsupported(ignoreUnsupported bool) TreePrinterOption {
	return func(opts *TreePrinterOptions) {
		opts.IgnoreUnsupported = ignoreUnsupported
	}
}

// PrinterOptionPadding option specification.
func PrinterOptionPadding(padding uint) TreePrinterOption {
	return func(opts *TreePrinterOptions) {
		opts.Padding = padding
	}
}

// PrinterOptionPrettyNames option specification.
func PrinterOptionPrettyNames(names map[string]string) TreePrinterOption {
	return func(opts *TreePrinterOptions) {
		opts.PrettyNames = names
	}
}

// PrinterOptionIgnoreNames option specification.
func PrinterOptionIgnoreNames(names []string) TreePrinterOption {
	return func(opts *TreePrinterOptions) {
		opts.IgnoreNames = names
	}
}
