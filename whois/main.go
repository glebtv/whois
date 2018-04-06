/*
 * Go module for domain whois
 * https://www.likexian.com/
 *
 * Copyright 2014-2018, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/glebtv/whois"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(fmt.Sprintf("usage:\n\t%s domain", os.Args[0]))
		os.Exit(1)
	}

	result := whois.Whois(os.Args[1])

	spew.Dump(result)
	os.Exit(0)
}
