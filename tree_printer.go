package afmt

import (
	"fmt"
	"reflect"
)

// TreePrinter definition.
type TreePrinter struct {
	//
	ignoreUnsupported bool
	padding           uint
	prettyNames       map[string]string
	ignoreNames       []string
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

// Print provides printing of structures or other ...
func (p *TreePrinter) Print(inspect interface{}) (string, error) {
	var level = []bool{}
	return p.inspectInterface(inspect, level, "")
}

func (p *TreePrinter) printPointer(inspect interface{}, lvl []bool) (string, error) {
	if p.isIgnored(reflect.TypeOf(inspect).Name()) {
		return "", nil
	}

	if reflect.ValueOf(inspect).Pointer() == 0 {
		return "", nil
	}

	return p.inspectInterface(reflect.Indirect(reflect.ValueOf(inspect)).Interface(), lvl, "")
}

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

	var ires string
	var level = append(lvl, false)

	if reflect.TypeOf(inspect).Kind() == reflect.Struct {
		val := reflect.ValueOf(inspect)

		for i := 0; i < val.NumField(); i++ {
			if !p.isIgnored(val.Type().Field(i).Name) {
				if i == val.NumField()-1 {
					level[len(level)-1] = true // set true if it is last item
				}

				if res, err := p.inspectInterface(val.Field(i).Interface(), level, p.getPrettyName(val.Type().Field(i).Name)); err != nil {
					return "", err
				} else {
					ires += res
				}
			}
		}
	}

	return structName + ires, nil
}

func (p *TreePrinter) printArray(inspect interface{}, lvl []bool, name string) (string, error) {
	if p.isIgnored(reflect.TypeOf(inspect).Name()) {
		return "", nil
	}

	arrayName := fmt.Sprintf("%s%s:\n", p.getPrefix(lvl), name)

	var ires string
	var level = append(lvl, false)

	if reflect.TypeOf(inspect).Kind() == reflect.Array || reflect.TypeOf(inspect).Kind() == reflect.Slice {
		val := reflect.ValueOf(inspect)

		for i := 0; i < val.Len(); i++ {
			if i == val.Len()-1 {
				level[len(level)-1] = true // set true if it is last item
			}

			if res, err := p.inspectInterface(val.Index(i).Interface(), level, ""); err != nil {
				return "", err
			} else {
				ires += res
			}
		}
	}

	return arrayName + ires, nil
}

func (p *TreePrinter) printString(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%s\n", p.getPrefix(lvl), reflect.ValueOf(inspect).String())
	}
	return fmt.Sprintf("%s%s: %s\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).String())
}

func (p *TreePrinter) printBoolean(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%t\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Bool())
	}
	return fmt.Sprintf("%s%s: %t\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Bool())

}

func (p *TreePrinter) printInteger(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%d\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Int())
	}
	return fmt.Sprintf("%s%s: %d\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Int())

}

func (p *TreePrinter) printFloat(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%f\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Float())
	}
	return fmt.Sprintf("%s%s: %f\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Float())

}

func (p *TreePrinter) printUnsignedInteger(inspect interface{}, lvl []bool, name string) string {
	if name == "" {
		return fmt.Sprintf("%s%d\n", p.getPrefix(lvl), reflect.ValueOf(inspect).Uint())
	}
	return fmt.Sprintf("%s%s: %d\n", p.getPrefix(lvl), name, reflect.ValueOf(inspect).Uint())

}

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

func (p *TreePrinter) getPrettyName(name string) string {
	if p.prettyNames[name] != "" {
		return p.prettyNames[name]
	}

	return name
}

func (p *TreePrinter) applyPadding(filler string) string {
	var fill string

	for i := 0; i < int(p.padding)-2; i++ {
		fill += filler
	}

	return fill
}

func (p *TreePrinter) isIgnored(name string) bool {
	for _, n := range p.ignoreNames {
		if name == n {
			return true
		}
	}

	return false
}
