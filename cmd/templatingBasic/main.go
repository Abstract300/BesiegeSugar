package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

const (
	pwd = "/home/abstract300/Programming/BesiegeSugar/templates/"
)

func main() {
	tmpls := filepath.Join(pwd, "*")
	t := template.Must(template.ParseGlob(tmpls))
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
}
