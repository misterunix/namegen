package main

import (
	"fmt"
	"os"
)

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Check err and exit if fatal is true or return err if fatal is false
// Print err to stderr
func CheckErr(err error, fatal bool) error {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if fatal {
			os.Exit(1)
		}
	}
	return err
}
