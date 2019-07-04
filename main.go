package main

import (
	"os"
	"text/template"
	"fmt"
	"strings"
	"gopkg.in/yaml.v2"
	"bytes"
)

type DockerImageRef struct {
	Repository  string
	Path        string
	Tag         string
}

func readConfig(source string) DockerImageRef {
	var config DockerImageRef

	fmt.Printf("%s", source)

	err := yaml.Unmarshal([]byte(source), &config)

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return config
}

func render(ref DockerImageRef) string {

	fullpath := strings.Join([]string{ref.Repository, ref.Path}, "/")
	return fmt.Sprintf("https://%v:%v", fullpath, ref.Tag)
}

func parse(source string) string {
    t, err := template.New("config").Funcs(template.FuncMap{
		"getenv": func(name string, fallback string) string {
			val := os.Getenv(name)
			if (val == "") {
				return fallback
			}
			
			return val
		},
  	}).Parse(source)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, []string{})
	if err != nil {
		panic(err)
	}

	return tpl.String()
}

func main() {
	os.Setenv("NAME", "gopher")
	raw := `
repository: "{{ getenv "FIRSTNAME" "cos" }}"
path: "kolemp/gct"
tag: a123
`
	parsed := parse(raw)
	ref := readConfig(parsed)

	fmt.Printf(render(ref))
}