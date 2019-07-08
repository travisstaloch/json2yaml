package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
)

//go:generate go run gen.go

type converter func([]byte) ([]byte, error)

const (
	usageDefault = `
Usage: json2yaml [OPTION?] [FILENAME*]
Convert between json and yaml formats.
Input(s) supplied by stdin or as FILENAME(s).
If format is given, all files will be converted to format.
Otherwise .json files are converted to yaml and .yaml to json.

OPTION:
	-f --format {yaml (default), json} : examples: --f=yaml -format json

examples:
  $ cat 1.json | json2yaml
  $ cat 1.yaml | json2yaml -format json | tee file.json
  $ echo {"a": 1, "b": 2} | json2yaml -f yaml
  $ printf "a: 1\\nb: 2" | json2yaml --f json
  $ json2yaml file.json 2.yaml | tee file.yaml
  $ json2yaml -f=json file.yaml 2.json
`
)

func main() {
	const (
		fmtU  = "Output format (json or yaml)"
		short = " (shorthand)"
	)
	var format string
	flag.StringVar(&format, "format", "<empty>", fmtU)
	flag.StringVar(&format, "f", "<empty>", fmtU+short)
	flag.Usage = func() {
		pkgusage := usage()
		if len(pkgusage) > 0 {
			fmt.Println(pkgusage)
		} else {
			fmt.Println(usageDefault)
		}
		os.Exit(0)
	}
	flag.Parse()

	conversions := map[string]converter{
		"json":    yaml.YAMLToJSON,
		"yaml":    yaml.JSONToYAML,
		"<empty>": nil,
	}

	inversions := map[string]string{
		"json": "yaml",
		"yaml": "json",
	}

	fn, formatOk := conversions[strings.ToLower(format)]
	if !formatOk {
		fmt.Fprintf(os.Stderr, "Unknown format '%s'\n", format)
		return
	}
	formatGiven := fn != nil

	var filenames = flag.Args()
	if len(filenames) > 0 {
		for _, filename := range filenames {
			// var fn2 converter
			bytes, readerr := os.Open(filename)
			check(readerr)

			if !formatGiven {
				ext := strings.ToLower(filepath.Ext(filename)[1:])
				inv, okInv := inversions[ext]
				checkBoolMsg(okInv, "Invalid file extension '%s'\n", ext)
				fn, _ = conversions[inv]
			}
			convert(bytes, fn)
		}
	} else {
		if fn == nil {
			fn = yaml.JSONToYAML
		}
		convert(os.Stdin, fn)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkMsg(e error, f string, msg ...interface{}) {
	if e != nil {
		fmt.Fprintf(os.Stderr, f, msg...)
		panic(e)
	}
}

func checkBoolMsg(b bool, f string, msg ...interface{}) {
	if !b {
		fmt.Fprintf(os.Stderr, f, msg...)
		os.Exit(1)
	}
}

func convert(f *os.File, fn converter) {
	var lines, readerr = ioutil.ReadAll(f)
	check(readerr)

	var r, converterr = fn(lines)
	check(converterr)
	os.Stdout.Write(r)
}
