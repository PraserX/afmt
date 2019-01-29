// Package afmt (Advanced formatter or Advanced fmt) implement some extensions
// for fmt package. The main feature is print to tree structure.
package afmt

import (
	"fmt"
	"reflect"
)

// TreePrinter definition.
type TreePrinter struct {
	// If ignore supported is true, then unimplemented types will be skipped.
	// Otherwise error is returned. Default value is true.
	ignoreUnsupported bool
	// Base padding definition for level. Default value is 4.
	padding uint
	// List of pretty names, which replace original values.
	prettyNames map[string]string
	// List of items, which are ignored.
	ignoreNames []string
}

// NewTreePrinter create new TreePrinter `instance`.
func NewTreePrinter(opts ...TreePrinterOption) *TreePrinter {
	var options = &TreePrinterOptions{
		IgnoreUnsupported: true,
		Padding:           4,
		PrettyNames:       map[string]string{},
		IgnoreNames:       []string{},
	}

	for _, opt := range opts {
		opt(options)
	}

	return &TreePrinter{
		ignoreUnsupported: options.IgnoreUnsupported,
		padding:           options.Padding,
		prettyNames:       options.PrettyNames,
		ignoreNames:       options.IgnoreNames,
	}
}

// Print method represents entry point for
func (p *TreePrinter) Print(inspect interface{}) (string, error) {
	var level = []bool{}
	return p.inspectInterface(inspect, level, "")
}

// printPointer transforms pointer to value and continue with interface
// inspection. If pointer is nil, then empty string is returned.
func (p *TreePrinter) printPointer(inspect interface{}, lvl []bool) (string, error) {
	if p.isIgnored(reflect.TypeOf(inspect).Name()) {
		return "", nil
	}

	if reflect.ValueOf(inspect).Pointer() == 0 {
		return "", nil
	}

	return p.inspectInterface(reflect.Indirect(reflect.ValueOf(inspect)).Interface(), lvl, "")
}

// printStruct provides extraction of structure items and continue with
// interface inspection. Extraction process also includes elimination of ignored
// values.
func (p *TreePrinter) printStruct(inspect interface{}, lvl []bool) (string, error) {
	if p.isIgnored(reflect.TypeOf(inspect).Name()) {
		return "", nil
	}

	var structName string

	if reflect.TypeOf(inspect).Name() != "" {
		structName = fmt.Sprintf("%s%s:\n", p.getPrefix(lvl), p.getPrettyName(reflect.TypeOf(inspect).Name()))
	} else {
		structName = fmt.Sprintf("%s%s\n", p.getPrefix(lvl), "■")
	}

	var err error
	var ires, res string
	var level = append(lvl, false)

	if reflect.TypeOf(inspect).Kind() == reflect.Struct {
		val := reflect.ValueOf(inspect)

		for i := 0; i < val.NumField(); i++ {
			if !p.isIgnored(val.Type().Field(i).Name) {
				if p.isLastNonIgnored(inspect, i) {
					level[len(level)-1] = true // set true if it is last item
				}

				if res, err = p.inspectInterface(val.Field(i).Interface(), level, p.getPrettyName(val.Type().Field(i).Name)); err != nil {
					return "", err
				}

				ires += res
			}
		}
	}

	return structName + ires, nil
}

// printArray provides extraction of array items and continue with interface
// inspection. Extraction process also includes elimination of ignored values.
func (p *TreePrinter) printArray(inspect interface{}, lvl []bool, name string) (string, error) {
	if p.isIgnored(reflect.TypeOf(inspect).Name()) {
		return "", nil
	}

	arrayName := fmt.Sprintf("%s%s:\n", p.getPrefix(lvl), name)

	var err error
	var ires, res string
	var level = append(lvl, false)

	if reflect.TypeOf(inspect).Kind() == reflect.Array || reflect.TypeOf(inspect).Kind() == reflect.Slice {
		val := reflect.ValueOf(inspect)

		for i := 0; i < val.Len(); i++ {
			if p.isLastNonIgnored(inspect, i) {
				level[len(level)-1] = true // set true if it is last item
			}

			if res, err = p.inspectInterface(val.Index(i).Interface(), level, ""); err != nil {
				return "", err
			}

			ires += res
		}
	}

	return arrayName + ires, nil
}

// printString returns finalized string output of string.
func (p *TreePrinter) printString(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%s\n", p.getPrefix(lvl), reflect.ValueOf(inspect).String())
	}
	return fmt.Sprintf("%s%s: %s\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).String())
}

// printBoolean returns finalized string output of boolean.
func (p *TreePrinter) printBoolean(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%t\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Bool())
	}
	return fmt.Sprintf("%s%s: %t\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Bool())

}

// printInteger returns finalized string output of integer.
func (p *TreePrinter) printInteger(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%d\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Int())
	}
	return fmt.Sprintf("%s%s: %d\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Int())

}

// printFloat returns finalized string output of float.
func (p *TreePrinter) printFloat(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%f\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Float())
	}
	return fmt.Sprintf("%s%s: %f\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Float())

}

// printUnsignedInteger returns finalized string output of unsigned interger.
func (p *TreePrinter) printUnsignedInteger(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%d\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Uint())
	}
	return fmt.Sprintf("%s%s: %d\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Uint())

}

// inspectInterface is an elementary method which provides interface inspection.
// Based on type of object on input, it calls specialized method for deep
// inspection of type and value. This method is called recursively every time,
// when inspection is needed.
func (p TreePrinter) inspectInterface(inspect interface{}, level []bool, name string) (string, error) {
	var err error
	var res string

	val := reflect.ValueOf(inspect)

	switch val.Kind() {
	case reflect.Ptr:
		res, err = p.printPointer(inspect, level)
	case reflect.Struct:
		res, err = p.printStruct(inspect, level)
	case reflect.Array, reflect.Slice:
		res, err = p.printArray(inspect, level, name)
	case reflect.String:
		res = p.printString(inspect, level, name)
	case reflect.Bool:
		res = p.printBoolean(inspect, level, name)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res = p.printInteger(inspect, level, name)
	case reflect.Float32, reflect.Float64:
		res = p.printFloat(inspect, level, name)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res = p.printUnsignedInteger(inspect, level, name)
	default:
		if p.ignoreUnsupported == false {
			return "", fmt.Errorf("Type %s", reflect.TypeOf(val))
		}
	}

	if err != nil {
		return "", err
	}

	return res, nil
}

// getPrefix returns string prefix with applied paddings. These paddings are
// applied according to level.
func (p *TreePrinter) getPrefix(lvl []bool) string {
	var levelPrefix string
	var level = len(lvl)

	for i := 0; i < level; i++ {
		if level == 1 && lvl[i] == true {
			levelPrefix += fmt.Sprintf("└%s ", p.applyPadding("─"))
		} else if level == 1 && lvl[i] == false {
			levelPrefix += fmt.Sprintf("├%s ", p.applyPadding("─"))
		} else if i+1 == level && lvl[i] == false {
			levelPrefix += fmt.Sprintf("├%s ", p.applyPadding("─"))
		} else if i+1 == level && lvl[i] == true {
			levelPrefix += fmt.Sprintf("└%s ", p.applyPadding("─"))
		} else if lvl[i] == true {
			levelPrefix += fmt.Sprintf(" %s ", p.applyPadding(" "))
		} else {
			levelPrefix += fmt.Sprintf("│%s ", p.applyPadding(" "))
		}
	}

	return levelPrefix
}

// getPrettyName returns pretty name based on specification. Pretty names are
// defined during TreePrinter initialization.
func (p *TreePrinter) getPrettyName(name string) string {
	if p.prettyNames[name] != "" {
		return p.prettyNames[name]
	}

	return name
}

// applyPadding is a function for support of variable padding. It returns string
// which contains crafted value based on specified filler.
func (p *TreePrinter) applyPadding(filler string) string {
	var fill string

	for i := 0; i < int(p.padding)-2; i++ {
		fill += filler
	}

	return fill
}

// isIgnored return true value if name of item is registered in ignore list.
// Otherwire false is returned.
func (p *TreePrinter) isIgnored(name string) bool {
	for _, n := range p.ignoreNames {
		if name == n {
			return true
		}
	}

	return false
}

// isLastNonIgnored return true if value (specially some interface) is last
// non-ignored item in list or structure, otherwire returns false. It is
// required, that inspected value has to be Struct, Array or Slice. If not,
// panic is fired.
func (p *TreePrinter) isLastNonIgnored(inspect interface{}, index int) bool {
	val := reflect.ValueOf(inspect)

	if reflect.TypeOf(inspect).Kind() == reflect.Struct {
		for i := index + 1; i < val.NumField(); i++ {
			if !p.isIgnored(reflect.TypeOf(val.Field(i).Interface()).Name()) {
				return false
			}
		}
	} else if reflect.TypeOf(inspect).Kind() == reflect.Array || reflect.TypeOf(inspect).Kind() == reflect.Slice {
		for i := index + 1; i < val.Len(); i++ {
			if !p.isIgnored(reflect.TypeOf(val.Index(i).Interface()).Name()) {
				return false
			}
		}
	} else {
		panic("Inspected value is not Struct nor Array nor Slice")
	}

	return true
}
