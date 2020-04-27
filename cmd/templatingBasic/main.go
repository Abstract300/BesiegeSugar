package main

import (
	"github.com/abstract300/besiegesugar"
)

const (
	pwd = "/home/abstract300/Programming/BesiegeSugar/templates/"
)

func main() {
	post := besiegesugar.ApplyConfig("001-example.cfg")
	post.MakePost()
}
