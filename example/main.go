package main

import (
	"fmt"
	"log"

	"github.com/AngusLu/argparse"
)

func main() {
	fmt.Println("argparse example")
	opts := []string{
		// "-f", "/config.yml",
		// "-v",
		// "--port", "3000",
		// "--float", "1.23",
		// "-n", "name!!",
		// "-i", "item1",
		// "-i", "item2",
		// "-h",
	}

	p := argparse.New()

	var configPath string
	var versionFlag bool
	var port int
	var floatVal float64
	var name *string
	var sarr []string
	var missReq int
	_ = missReq

	p.BoolVar(&versionFlag, false, "v", "version", "show version", nil)
	p.StringVar(&configPath, "", "f", "config", "", nil)
	p.IntVar(&port, 0, "p", "port", "http listen port", nil)
	p.FloatVar(&floatVal, 1.2, "ff", "float", "float value", nil)
	p.StringSliceVar(&sarr, []string{}, "i", "item", "item list", nil)
	// uncomment to test required arg missing
	// p.IntVar(&missReq, -1, "m", "miss", "require value", &argparse.Option{Require: true})
	p.Help("h", "help")

	name = p.String("", "n", "name", "show name", nil)

	err := p.Parse(opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"Config: %s\nVersionFlag: %v\nPort: %d\nFloatVal: %f\nName: %s\nItems: %v\n",
		configPath,
		versionFlag,
		port,
		floatVal,
		*name,
		sarr,
	)
}
