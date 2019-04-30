// Package afmt (Advanced formatter or Advanced fmt) implement some extensions
// for fmt package. The main feature is print to tree structure.
package afmt

import (
	"strings"
	"unicode"
)

// ColPrinter definition.
type ColPrinter struct{}

// NewColPrinter allocates new column printer object.
func NewColPrinter() *ColPrinter {
	printer := &ColPrinter{}
	return printer
}

// Print adding spaces between operands when neither is a string. Newline is
// always added in order to respect column maximum characters (lineLimit). It
// return string with maximum length of line based on lineLimit.
func (c *ColPrinter) Print(lineLimit int, a ...interface{}) (out string) {
	var buf string

	for _, text := range a {
		for _, s := range text.(string) {
			if !(len(buf) == 0 && unicode.IsSpace(s)) {
				buf += string(s)
			}

			if len(buf) >= lineLimit {
				if space := strings.LastIndexFunc(buf, unicode.IsSpace); space != -1 {
					out += buf[:space] + "\n"
					buf = buf[space+1:]
				} else {
					out += buf + "\n"
					buf = ""
				}
			}
		}
		buf += " "
	}

	out += buf[:len(buf)-1]
	return out
}
