package afmt

import (
	"fmt"
	"reflect"
)

// Printer definition.
type Printer struct {
	//
	ignoreUnsupported bool
	padding           uint
	prettyNames       map[string]string
	ignoreNames       []string
}

// NewPrinter create new Printer `instance`.
func NewPrinter(opts ...PrinterOption) *Printer {
	var options = &PrinterOptions{
		IgnoreUnsupported: true,
		Padding:           4,
		PrettyNames:       map[string]string{},
		IgnoreNames:       []string{},
	}

	for _, opt := range opts {
		opt(options)
	}

	return &Printer{
		ignoreUnsupported: options.IgnoreUnsupported,
		padding:           options.Padding,
		prettyNames:       options.PrettyNames,
		ignoreNames:       options.IgnoreNames,
	}
}

// Tree provides printing of structures or other ...
func (p *Printer) Tree(inspect interface{}) error {
	var level = []bool{}
	return p.inspectInterface(inspect, level, "")
}

func (p *Printer) printStruct(inspect interface{}, lvl []bool) {
	if p.isIgnored(reflect.TypeOf(inspect).Name()) {
		return
	}

	if reflect.TypeOf(inspect).Name() != "" {
		fmt.Printf("%s%s:\n", p.getPrefix(lvl), p.getPrettyName(reflect.TypeOf(inspect).Name()))
	} else {
		fmt.Printf("%s%s\n", p.getPrefix(lvl), "■")
	}

	var level = append(lvl, false)

	if reflect.TypeOf(inspect).Kind() == reflect.Struct {
		val := reflect.ValueOf(inspect)

		for i := 0; i < val.NumField(); i++ {
			if !p.isIgnored(val.Type().Field(i).Name) {
				if i == val.NumField()-1 {
					level[len(level)-1] = true // set true if it is last item
				}

				p.inspectInterface(val.Field(i).Interface(), level, p.getPrettyName(val.Type().Field(i).Name))
			}
		}
	}
}

func (p *Printer) printArray(inspect interface{}, lvl []bool, name string) {
	if p.isIgnored(reflect.TypeOf(inspect).Name()) {
		return
	}

	fmt.Printf("%s%s:\n", p.getPrefix(lvl), name)

	var level = append(lvl, false)

	if reflect.TypeOf(inspect).Kind() == reflect.Array || reflect.TypeOf(inspect).Kind() == reflect.Slice {
		val := reflect.ValueOf(inspect)

		for i := 0; i < val.Len(); i++ {
			if i == val.Len()-1 {
				level[len(level)-1] = true // set true if it is last item
			}

			p.inspectInterface(val.Index(i).Interface(), level, "")
		}
	}
}

func (p *Printer) printString(inspect interface{}, lvl []bool, name string) {
	if name == "" {
		fmt.Printf("%s%s\n", p.getPrefix(lvl), reflect.ValueOf(inspect).String())
	} else {
		fmt.Printf("%s%s: %s\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).String())
	}
}

func (p *Printer) printBoolean(inspect interface{}, lvl []bool, name string) {
	if name == "" {
		fmt.Printf("%s%t\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Bool())
	} else {
		fmt.Printf("%s%s: %t\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Bool())
	}
}

func (p *Printer) printInteger(inspect interface{}, lvl []bool, name string) {
	if name == "" {
		fmt.Printf("%s%d\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Int())
	} else {
		fmt.Printf("%s%s: %d\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Int())
	}
}

func (p *Printer) printFloat(inspect interface{}, lvl []bool, name string) {
	if name == "" {
		fmt.Printf("%s%f\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Float())
	} else {
		fmt.Printf("%s%s: %f\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Float())
	}
}

func (p *Printer) printUnsignedInteger(inspect interface{}, lvl []bool, name string) {
	if name == "" {
		fmt.Printf("%s%d\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Uint())
	} else {
		fmt.Printf("%s%s: %d\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Uint())
	}
}

func (p Printer) inspectInterface(inspect interface{}, level []bool, name string) error {
	val := reflect.ValueOf(inspect)

	switch val.Kind() {
	case reflect.Struct:
		p.printStruct(inspect, level)

	case reflect.Array, reflect.Slice:
		p.printArray(inspect, level, name)

	case reflect.String:
		p.printString(inspect, level, name)

	case reflect.Bool:
		p.printBoolean(inspect, level, name)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p.printInteger(inspect, level, name)

	case reflect.Float32, reflect.Float64:
		p.printFloat(inspect, level, name)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		p.printUnsignedInteger(inspect, level, name)

	default:
		if p.ignoreUnsupported == false {
			return fmt.Errorf("Type %s", reflect.TypeOf(val))
		}
	}

	return nil
}

func (p *Printer) getPrefix(lvl []bool) string {
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

func (p *Printer) getPrettyName(name string) string {
	if p.prettyNames[name] != "" {
		return p.prettyNames[name]
	}

	return name
}

func (p *Printer) applyPadding(filler string) string {
	var fill string

	for i := 0; i < int(p.padding)-2; i++ {
		fill += filler
	}

	return fill
}

func (p *Printer) isIgnored(name string) bool {
	for _, n := range p.ignoreNames {
		if name == n {
			return true
		}
	}

	return false
}
