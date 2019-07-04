package main

import (
	// "os"
	// "text/template"
	"fmt"
	"strings"
)

type DockerImageRef struct {
	Repository  string
	Path        string
	Tag         string
}

func render(ref DockerImageRef) string {
	fullpath := strings.Join([]string{ref.Repository, ref.Path}, "/")
	return fmt.Sprintf("https://%v:%v", fullpath, ref.Tag)
}


func main() {
    // t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	// if err != nil {
	// 	panic(err)
	// }
	// err = t.Execute(os.Stdout, td)
	// if err != nil {
	// 	panic(err)
	// }
	ref := DockerImageRef{"hub", "kolemp", "12"}
	fmt.Printf(render(ref))
}