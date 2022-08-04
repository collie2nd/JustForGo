package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.PanicOnError)
		goCmd.StringVar(&name, "name", "go", "help")
		goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ContinueOnError)
		phpCmd.StringVar(&name, "n", "php", "help")
		phpCmd.Parse(args[1:])
	}
	log.Println("name:", name)
}
