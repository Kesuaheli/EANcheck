package main

import (
	"eancheck/gs1"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	help = flag.Bool("help", false, "Prints usage and exit.")
	//mode = flag.String("mode", "unparse", "Sets the mode\n.")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		flag.PrintDefaults()
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		fmt.Printf("Expected EAN code:\n%s [--flags] <EAN-Code>\n", os.Args[0])
		os.Exit(1)
	}

	code := flag.Arg(0)

	if len(code) != 13 {
		fmt.Printf("EAN code must be 13 digits!")
		os.Exit(1)
	}

	if _, err := strconv.Atoi(code); err != nil {
		fmt.Printf("EAN code must only contain digits!")
		os.Exit(1)
	}

	log.Printf("Code format is valid: %s", code)

	prefix, rest := gs1.ParseEAN13(code)

	log.Printf("Country: %s\n", prefix)
	log.Printf("Rest:    %s\n", rest)
}
