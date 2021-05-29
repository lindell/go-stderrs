package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

type errorGen struct {
	Name           string `yaml:"name"`
	Description    string `yaml:"description"`
	HTTPStatusCode int    `yaml:"httpStatusCode"`
	Temporary      bool   `yaml:"temporary"`
}

type genFile struct {
	Errors []errorGen `yaml:"errors"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("./source.yaml")
	must(err)

	defer f.Close()

	var genFile genFile
	err = yaml.NewDecoder(f).Decode(&genFile)
	must(err)

	errTmpl, err := template.ParseFiles("./generate/error.tmpl")
	must(err)

	readmeTmpl, err := template.ParseFiles("./generate/README.tmpl.md")
	must(err)

	readmeF, err := os.Create("README.md")
	must(err)
	defer readmeF.Close()
	must(readmeTmpl.Execute(readmeF, genFile))

	must(os.RemoveAll("stderrs"))
	must(os.Mkdir("stderrs", 0744))

	for _, e := range genFile.Errors {
		f, err := os.Create(filepath.Join("./stderrs", fmt.Sprintf("%s.go", strings.ToLower(e.Name))))
		must(err)
		defer f.Close()
		must(errTmpl.Execute(f, e))
	}
}
