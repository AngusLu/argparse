package argparse

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Parser struct {
	args     []*arg
	showHelp bool
	parsed   bool
}

type Option struct {
	Require bool
}

type arg struct {
	sname string
	lname string
	// value item count
	size         int
	value        interface{}
	defaultValue interface{}
	description  string
	unique       bool
	parsed       bool
	opts         *Option
}

func New() *Parser {
	return &Parser{
		args: make([]*arg, 0),
	}
}

func (p *Parser) addArg(a *arg) error {
	for _, v := range p.args {
		if v.sname == a.sname || v.lname == a.lname {
			return errors.New("option name dup")
		}
	}
	p.args = append(p.args, a)

	return nil
}

func (p *Parser) typeVar(i interface{}, defVal interface{}, short, long, description string, size int, unique bool, opts *Option) {
	t := reflect.ValueOf(i)
	if t.Kind() != reflect.Ptr {
		panic(errors.New("var type not ptr"))
	}
	if short == "" && long == "" {
		panic(errors.New("short name and long name is empty"))
	}

	a := &arg{
		sname:        short,
		lname:        long,
		value:        i,
		defaultValue: defVal,
		description:  description,
		size:         size,
		unique:       unique,
		opts:         opts,
	}

	if err := p.addArg(a); err != nil {
		panic(fmt.Errorf("unable to add String: %s\n", err))
	}
}

func (p *Parser) Help(short, long string) {
	p.typeVar(&p.showHelp, false, short, long, "show usage help", 0, true, nil)
}

func (p *Parser) printHelp() {
	sb := &strings.Builder{}
	sb.WriteString("Usage:\n")
	for _, v := range p.args {
		if _, err := sb.WriteString(fmt.Sprintf("\t%s\t%s\t\t%s\n", v.name(), v.getType(), v.description)); err != nil {
			panic(err)
		}
	}
	fmt.Println(sb.String())
	// if flag.Lookup("test.v") == nil {
	os.Exit(2)
	// }
}

func (p *Parser) Parse(a []string) error {
	copyArg := make([]string, len(a))
	copy(copyArg, a)

	err := p.parse(&copyArg)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) parse(args *[]string) error {
	if p.parsed {
		return nil
	}
	if len(p.args) < 1 {
		return nil
	}

	err := p.parseArguemtns(args)
	if err != nil {
		return err
	}

	p.parsed = true

	if p.showHelp == true {
		p.printHelp()
	}

	p.setDefaultValue()

	return p.checkRequired()
}

func (p *Parser) setDefaultValue() {
	for _, v := range p.args {
		// fmt.Printf("show %s , parsed: %v, defVal: %v\n", v.name(), v.parsed, v.defaultValue)
		if !v.parsed {
			t := reflect.ValueOf(v.value)
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			t.Set(reflect.ValueOf(v.defaultValue))
			// *v.value = v.defaultValue
		}
	}
}

func (p *Parser) checkRequired() (err error) {
	for _, v := range p.args {
		if v.opts != nil && v.opts.Require && !v.parsed {
			return fmt.Errorf("[%s] is required", v.name())
		}
	}
	return
}

func (p *Parser) parseArguemtns(args *[]string) error {
	n := len(p.args)

	for i := 0; i < n; i++ {
		nn := len(*args)
		oarg := p.args[i]

		for j := 0; j < nn; j++ {
			s := (*args)[j]
			if s == "" {
				continue
			}
			ct, err := oarg.check(s)
			if err != nil {
				return err
			}
			if ct > 0 {
				valpos := j + 1
				if len(*args) < valpos+oarg.size {
					return fmt.Errorf("no enough arguments for %s", oarg.name())
				}
				err := oarg.parse((*args)[valpos : valpos+oarg.size])
				if err != nil {
					return err
				}

				oarg.reduce(j, args)
				continue
			}
		}
		// if oarg.opts != nil && oarg.opts.Require && !oarg.parsed {
		//   return fmt.Errorf("[%s] is required", oarg.name())
		// }

	}

	return nil
}

func (p *Parser) String(defaultValue, short, long, description string, opts *Option) *string {
	var result string

	p.typeVar(&result, defaultValue, short, long, description, 1, true, opts)

	return &result
}
func (p *Parser) StringVar(i *string, defaultValue, short, long, description string, opts *Option) {
	p.typeVar(i, defaultValue, short, long, description, 1, true, opts)
}

func (p *Parser) Int(defaultValue int, short, long, description string, opts *Option) *int {
	var result int

	p.typeVar(&result, defaultValue, short, long, description, 1, true, opts)

	return &result
}
func (p *Parser) IntVar(i *int, defaultValue int, short, long, description string, opts *Option) {
	p.typeVar(i, defaultValue, short, long, description, 1, true, opts)
}

func (p *Parser) Bool(defaultValue bool, short, long, description string, opts *Option) *bool {
	var result bool

	p.typeVar(&result, defaultValue, short, long, description, 0, true, opts)

	return &result
}
func (p *Parser) BoolVar(i *bool, defaultValue bool, short, long, description string, opts *Option) {
	p.typeVar(i, defaultValue, short, long, description, 0, true, opts)
}

func (p *Parser) Float(defaultValue float64, short, long, description string, opts *Option) *float64 {
	var result float64

	p.typeVar(&result, defaultValue, short, long, description, 1, true, opts)

	return &result
}
func (p *Parser) FloatVar(i *float64, defaultValue float64, short, long, description string, opts *Option) {
	p.typeVar(i, defaultValue, short, long, description, 1, true, opts)
}

func (p *Parser) StringSlice(defaultValue []string, short, long, description string, opts *Option) *[]string {
	var result []string

	p.typeVar(&result, defaultValue, short, long, description, 1, false, opts)

	return &result
}
func (p *Parser) StringSliceVar(i *[]string, defaultValue []string, short, long, description string, opts *Option) {
	p.typeVar(i, defaultValue, short, long, description, 1, false, opts)
}

func (p *Parser) IntSlice(defaultValue []int, short, long, description string, opts *Option) *[]int {
	var result []int

	p.typeVar(&result, defaultValue, short, long, description, 1, false, opts)

	return &result
}
func (p *Parser) IntSliceVar(i *[]int, defaultValue []int, short, long, description string, opts *Option) {
	p.typeVar(i, defaultValue, short, long, description, 1, false, opts)
}

func (p *Parser) FloatSlice(defaultValue []float64, short, long, description string, opts *Option) *[]float64 {
	var result []float64

	p.typeVar(&result, defaultValue, short, long, description, 1, false, opts)

	return &result
}
func (p *Parser) FloatSliceVar(i *[]float64, defaultValue []float64, short, long, description string, opts *Option) {
	p.typeVar(i, defaultValue, short, long, description, 1, false, opts)
}
