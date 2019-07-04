package main

import (
	"os"
	"text/template"
	"fmt"
	"strings"
	"gopkg.in/yaml.v2"
	"bytes"
	"flag"
	"io/ioutil"
)

var configFileName string

type DockerImageRef struct {
	Repository  string
	Path        string
	Tag         string
}

func readConfig(source string) DockerImageRef {
	var config DockerImageRef

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

func read (filename string) string {
	file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }

	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return string(b)
}


func init() {
	flag.StringVar(&configFileName, "c", "", "config file path")
}

func main() {

	flag.Parse()

	if (len(flag.Args()) == 0) {
		fmt.Println("Yoou need to pass operation. Currently supported are: imageName")
		os.Exit(1)
	}

	if (configFileName == "") {
		raw := `
repository: "{{ getenv "REPO_hostname" }}"
path: "image/path"
tag: "imageTag"
`
		fmt.Printf("You need to pass yaml config file with structure similar to this: \n %s", raw)
		os.Exit(1)
	}
	raw := read(configFileName)
	parsed := parse(raw)
	ref := readConfig(parsed)


	operation := flag.Args()[0]
	if (operation == "imageName") {

		fmt.Printf(render(ref))
	}
}