package argparse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrArgTooMany = errors.New("arg too manu")
	ErrNoArg      = errors.New("no arg")
)

func (a *arg) check(s string) (count int, err error) {
	res := a.checkLongName(s)
	if res > 0 {
		return res, nil
	}
	return a.checkShortName(s)
}

func (a *arg) checkLongName(s string) int {
	if a.lname != "" {
		if len(s) > 2 && strings.HasPrefix(s, "--") && s[2] != '-' {
			if s[2:] == a.lname {
				return 1
			}
		}
	}
	return 0
}

func (a *arg) checkShortName(s string) (int, error) {
	if a.sname != "" {
		if len(s) > 1 && strings.HasPrefix(s, "-") && s[1] != '-' {
			if s[1:] == a.sname {
				return 1, nil
			}
		}
	}
	return 0, nil
}

func (a *arg) reduce(pos int, args *[]string) {
	a.reduceLongName(pos, args)
	a.reduceShortName(pos, args)
}

func (a *arg) reduceLongName(pos int, args *[]string) {
	s := (*args)[pos]
	if a.lname == "" {
		return
	}

	if res := a.checkLongName(s); res == 0 {
		return
	}

	for i := pos; i <= pos+a.size; i++ {
		(*args)[i] = ""
	}
}

func (a *arg) reduceShortName(pos int, args *[]string) {
	s := (*args)[pos]
	if a.sname == "" {
		return
	}

	if res, err := a.checkShortName(s); res == 0 || err != nil {
		return
	}

	for i := pos; i <= pos+a.size; i++ {
		(*args)[i] = ""
	}
}

func (a *arg) name() string {
	if a.sname == "" {
		return fmt.Sprintf("--%s", a.lname)
	} else if a.lname == "" {
		return fmt.Sprintf("-%s", a.sname)
	} else {
		return fmt.Sprintf("-%s | --%s", a.sname, a.lname)
	}
}

func (a *arg) getType() string {
	switch a.value.(type) {
	case *string:
		return "string"
	case *int:
		return "int"
	case *bool:
		return "bool"
	case *float64:
		return "float"
	case *[]string:
		return "[]string"
	case *[]int:
		return "[]int"
	case *[]float64:
		return "[]float"
	default:
		return "not support"
	}
}

func (a *arg) parse(args []string) error {
	if a.unique && a.parsed {
		return fmt.Errorf("[%s] can only parsent once\n", a.name())
	}
	return a.parseType(args)
}

func (a *arg) parseType(args []string) error {
	var err error
	switch a.value.(type) {
	case *string:
		return a.parseString(args)
	case *int:
		return a.parseInt(args)
	case *bool:
		return a.parseBool(args)
	case *float64:
		return a.parseFloat(args)
	case *[]string:
		return a.parseStringSlice(args)
	case *[]int:
		return a.parseIntSlice(args)
	case *[]float64:
		return a.parseFloatSlice(args)
	default:
		err = fmt.Errorf("unsupport type [%t]", a.value)
	}
	return err
}

func (a *arg) parseString(args []string) error {
	var err error

	if len(args) > 1 {
		return ErrArgTooMany
	}
	if len(args) == 0 {
		return ErrNoArg
	}

	*((a.value).(*string)) = args[0]
	a.parsed = true

	return err
}

func (a *arg) parseStringSlice(args []string) (err error) {
	if len(args) == 0 {
		return ErrNoArg
	}

	*((a.value).(*[]string)) = append(*((a.value).(*[]string)), args...)
	a.parsed = true

	return
}

func (a *arg) parseInt(args []string) error {
	var err error

	if len(args) > 1 {
		return ErrArgTooMany
	}
	if len(args) == 0 {
		return ErrNoArg
	}

	if i, err := strconv.Atoi(args[0]); err == nil {
		*((a.value).(*int)) = i
		a.parsed = true
	} else {
		return fmt.Errorf("[%s] bad int value (%v)", a.name(), args[0])
	}

	return err
}

func (a *arg) parseIntSlice(args []string) (err error) {
	if len(args) == 0 {
		return ErrNoArg
	}

	for _, v := range args {
		if i, err := strconv.Atoi(v); err == nil {
			*((a.value).(*[]int)) = append(*((a.value).(*[]int)), i)
		} else {
			return fmt.Errorf("[%s] bad int value (%v)", a.name(), v)
		}
	}

	a.parsed = true

	return
}

func (a *arg) parseBool(args []string) error {
	*a.value.(*bool) = true
	a.parsed = true
	return nil
}

func (a *arg) parseFloat(args []string) (err error) {
	if len(args) > 1 {
		return ErrArgTooMany
	}
	if len(args) == 0 {
		return ErrNoArg
	}

	if f, err := strconv.ParseFloat(args[0], 64); err == nil {
		*a.value.(*float64) = f
		a.parsed = true
	} else {
		return fmt.Errorf("[%s] bad float value (%v)", a.name(), args[0])
	}
	return
}

func (a *arg) parseFloatSlice(args []string) (err error) {
	if len(args) == 0 {
		return ErrNoArg
	}

	for _, v := range args {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			*((a.value).(*[]float64)) = append(*((a.value).(*[]float64)), f)
		} else {
			return fmt.Errorf("[%s] bad float value (%v)", a.name(), v)
		}
	}
	a.parsed = true

	return
}
