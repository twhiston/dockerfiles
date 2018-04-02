package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type ImageDef struct {
	Name      string
	Tag       string
	Registry  string
	Namespace string
}

func main() {

	fmt.Println("--> Rendering codeship files for current image folders")
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(wd)
	if err != nil {
		log.Fatal(err)
	}

	dataset := make(map[string]interface{})
	subdirs := make([]ImageDef, 0)
	for _, f := range files {
		if f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
			fmt.Println(f.Name())
			subdirs = append(subdirs, ImageDef{Name: f.Name(), Namespace: "tomwhiston", Registry: "registry.twhiston.cloud"})
		}
	}
	dataset["services"] = subdirs

	tmpl := template.Must(template.ParseGlob(".util/ci-render/tmpl/*.tmpl"))

	render("services", dataset, tmpl)
	render("steps", dataset, tmpl)
}

func render(tName string, dataset map[string]interface{}, tmpl *template.Template) {
	f, err := os.Create("codeship-"+tName + ".yml")
	handleError(err)
	defer func(r *os.File) {
		err = r.Close()
		handleError(err)
	}(f)

	err = tmpl.ExecuteTemplate(f, tName+".tmpl", &dataset)
	handleError(err)
}

//HandleError simply logs the error and hard quits the app
func handleError(err error) {
	if err != nil {
		if err.Error() != "EOF" {
			log.Fatalln(err)
		}
	}
}
