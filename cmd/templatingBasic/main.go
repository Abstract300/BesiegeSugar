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
	tmpls := filepath.Join(pwd, "*")
	post := besiegesugar.Post{
		Title:   "Test title",
		Author:  "Abstract300",
		Content: "this is a test content.",
	}
	t := template.Must(template.ParseGlob(tmpls))
	err := t.Execute(os.Stdout, post)
	if err != nil {
		log.Fatal(err)
	}
}
