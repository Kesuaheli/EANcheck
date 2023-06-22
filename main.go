package main

import (
	"eancheck/ean"
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
		fmt.Printf("EAN code must be 13 digits!\n")
		os.Exit(1)
	}

	if _, err := strconv.Atoi(code); err != nil {
		fmt.Printf("EAN code must only contain digits!\n")
		os.Exit(1)
	}

	log.Printf("Code format is valid: %s\n", code)

	ean13, err := ean.ParseEAN13(code)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Country: %s\n", ean13.Prefix)
	log.Printf("Product: %s\n", ean13.Data)
	log.Printf("Parity:  %d\n", ean13.Parity)
}
