// +build ignore

// This program generates usage.go. It can be invoked by running
// go generate
package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {

	s, err := ioutil.ReadFile("README.md")
	check(err)
	u := bytes.Split(s, []byte("## Usage"))[1]
	f, err := os.Create("usage.go")
	check(err)
	defer f.Close()
	packageTemplate.Execute(f, struct{ S string }{S: string(u)})
}

var packageTemplate = template.Must(template.New("").Parse(`
package main

const(
	usage = ` + "`" + `{{ .S }}` + "`" + `
)
`))

func check(e error) {
	if e != nil {
		panic(e)
	}
}
