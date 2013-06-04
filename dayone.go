package main

import (
	"fmt"
	"os"
	"./dayone"
)

const version = "0.0.2"

func main() {
	argv := os.Args
	argc := len(argv)

	if argc < 2 {
		usage()
	}

	text := argv[1]
	entry := dayone.NewEntry(text)

	var dir string
	if argc == 3 {
		dir = argv[2]
	} else {
		dir = os.ExpandEnv("$HOME/Dropbox/Apps/Day One/Journal.dayone/entries")
	}

	entry.WriteIn(dir)
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: godayone text [dir]
Version %s
`, version)
	os.Exit(-1)
}
