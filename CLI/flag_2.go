package main

import (
	"flag"
	"log"
)

type Name string

func (s *Name) Set(val string) error {
	*s = Name("xxx" + val)
	return nil
}

func (s *Name) String() string {
	return string("yyy" + *s)
}

func main() {
	var name Name
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.PanicOnError)
		goCmd.Var(&name, "name", "go")
		goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ContinueOnError)
		phpCmd.Var(&name, "n", "php")
		phpCmd.Parse(args[1:])
	}
	log.Println("name:", name)
}
