package main

import (
	"fmt"
	"github.com/armhold/z2ll"
	"github.com/pkg/errors"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		die(errors.New("Usage: z2ll zipcode"))
	}

	zipCode := os.Args[1]

	l, err := z2ll.NewLocator()
	if err != nil {
		die(err)
	}

	zc, err := l.LookupZip(zipCode)
	if err != nil {
		die(err)
	}

	fmt.Printf("%f,%f\n", zc.Lat, zc.Long)
}

func die(e error) {
	fmt.Fprintf(os.Stderr, "%s\n", e)
	os.Exit(1)
}
