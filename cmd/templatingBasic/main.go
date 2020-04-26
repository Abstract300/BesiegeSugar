package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/abstract300/besiegesugar"
)

const (
	pwd = "/home/abstract300/Programming/BesiegeSugar/templates/"
)

func main() {
	post := besiegesugar.ApplyConfig("001-example.cfg")
	tmpls := filepath.Join(pwd, "*")
	t := template.Must(template.ParseGlob(tmpls))
	err := t.Execute(os.Stdout, post)
	if err != nil {
		log.Fatal(err)
	}
}
