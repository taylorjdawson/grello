package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

// Flags
var (
	_flag string
)

func main() {
	flag.Parse()
	flag.Visit(func(f *flag.Flag) { _flag = f.Shorthand })

	switch _flag {
	case "lb":
		fmt.Println("list boards")
	case "cb":
		fmt.Println("create board")
	case "sb":
		fmt.Println("sel board")
	case "db":
		fmt.Println("delete board")
	case "rb":
		fmt.Println("rename board")
	case "lc":
		fmt.Println("list cards")
	case "cc":
		fmt.Println("create card")
	case "vc":
		fmt.Println("view card")
	case "ec":
		fmt.Println("edit card")
	case "mc":
		fmt.Println("move card")
	case "rc":
		fmt.Println("rename card")
	case "ac":
		fmt.Println("archive  card")
	case "dc":
		fmt.Println("delete card")
	}
}

func init() {
	flag.BoolP("list-boards", "lb", true, "Lists all boards")
	//flag.StringVarP(&grello_flgg, "list-boards", "lb", "Lists all boards")
}

// If no param supplied
func emptyInput(param string) {
	if param == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

// printUsage is a custom function we created to print usage for our CLI app
func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
