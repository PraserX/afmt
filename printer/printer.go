package printer

import (
	"fmt"
	"reflect"
	// "github.com/kr/pretty"
)

// Printer definition.
type Printer struct {
	ignoreUnsupported bool
	padding           uint
}

// NewPrinter create new Printer `instance`.
func NewPrinter(opts ...Option) *Printer {
	var options = &Options{
		IgnoreUnsupported: true,
		Padding:           4,
	}

	for _, opt := range opts {
		opt(options)
	}

	return &Printer{
		ignoreUnsupported: options.IgnoreUnsupported,
		padding:           options.Padding,
	}
}

// Tree provides printing of structures or other ...
func (p *Printer) Tree(inspect interface{}) error {
	return nil
}

func (p *Printer) printTreeStruct(inspect interface{}) error {

	if reflect.TypeOf(inspect).Kind() == reflect.Struct {

		val := reflect.ValueOf(inspect)

		for i := 0; i < val.NumField(); i++ {
			switch val.Field(i).Kind() {
			case reflect.Bool:
				p.printTreeBoolean(val.Field(i))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				p.printTreeInteger(val.Field(i))
			case reflect.Float32, reflect.Float64:
				p.printTreeFloat(val.Field(i))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				p.printTreeUnsignedInteger(val.Field(i))
			case reflect.Array, reflect.Slice:
				p.printTreeArray(val.Field(i))
			case reflect.Struct:
				p.printTreeStruct(val.Field(i))
			default:
				if p.ignoreUnsupported == false {
					return fmt.Errorf("Type %s", reflect.TypeOf(val.Field(i)))
				}
			}
		}

		return nil
	}

	return fmt.Errorf("Value is not a structure")
}

func (p *Printer) printTreeArray(inspect interface{}) {

}

func (p *Printer) printTreeString(inspect interface{}) {

}

func (p *Printer) printTreeBoolean(inspect interface{}) {

}

func (p *Printer) printTreeInteger(inspect interface{}) {

}

func (p *Printer) printTreeFloat(inspect interface{}) {

}

func (p *Printer) printTreeUnsignedInteger(inspect interface{}) {

}

// if reflect.ValueOf(inspect).Kind() == reflect.Struct {

// 	t := reflect.TypeOf(inspect).Name()
// 	query := fmt.Sprintf("insert into %s values(", t)
// 	v := reflect.ValueOf(inspect)
// 	for i := 0; i < v.NumField(); i++ {
// 		switch v.Field(i).Kind() {
// 		case reflect.Int:
// 			if i == 0 {
// 				query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
// 			} else {
// 				query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
// 			}
// 		case reflect.String:
// 			if i == 0 {
// 				query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
// 			} else {
// 				query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
// 			}
// 		case reflect.Struct:
// 			fmt.Println("struct")
// 		default:
// 			fmt.Println("Unsupported type")
// 			return
// 		}
// 	}
// 	query = fmt.Sprintf("%s)", query)
// 	fmt.Println(query)
// 	return

// }
// fmt.Println("unsupported type")

// pretty.Println(inspect)
// pretty.Println(reflect.TypeOf(inspect))
// pretty.Println(reflect.ValueOf(inspect))

// fmt.Println(reflect.TypeOf(inspect))
// fmt.Println(reflect.ValueOf(inspect))

// fmt.Println(reflect.ValueOf(inspect))
