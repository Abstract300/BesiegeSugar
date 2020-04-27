package besiegesugar

import (
	"bufio"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	tmpls = "/home/abstract300/Programming/BesiegeSugar/templates/"
)

type Post struct {
	Filename    string
	Title       string
	Description string
	Keywords    string
	Created     string
	Updated     string
	Content     string
}

func LoadConfig(config string) []string {
	var data []string
	cfg, err := os.Open(config)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(cfg)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func ParseConfig(keyval []string) map[string]string {
	val := make(map[string]string, 0)

	for _, item := range keyval {
		valueIndex := strings.Index(item, "=")
		val[strings.ToLower(strings.Trim(item[:valueIndex-1], " "))] = strings.Trim(item[valueIndex+1:], " ")
	}

	return val
}

func ApplyConfig(file string) *Post {
	config := new(Post)

	cnf := ParseConfig(LoadConfig(file))

	config.Filename = cnf["filename"]
	config.Title = cnf["title"]
	config.Description = cnf["description"]
	config.Keywords = cnf["Keywords"]
	config.Created = cnf["created"]
	config.Updated = cnf["updated"]

	return config
}

func (p *Post) MakePost() {

	content, err := ioutil.ReadFile(p.Filename)
	if err != nil {
		log.Fatal(err)
	}
	p.Content = string(content)
	tmplate := template.Must(template.ParseGlob(filepath.Join(tmpls, "*.tmpl")))
	err = tmplate.ExecuteTemplate(os.Stdout, "body", p)
	if err != nil {
		log.Fatal(err)
	}
}
