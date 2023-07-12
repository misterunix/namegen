package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Reads a text file with only one column and adds it to a table
// used for firstnames, lastnames, etc. that do not contain frequencies
func readFileGenericAdd(f string, t string) {
	fmt.Printf("File: %s to table: %s\n", f, t)
	readFile, err := os.Open(f)
	_ = CheckErr(err, true)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	id := 0
	for fileScanner.Scan() {
		tmp := strings.TrimSpace(fileScanner.Text()) // remove spaces just in case
		namestruct := genericName{id, tmp}           // create a firstname struct
		sql := InsertIntoTable(t, namestruct)        // create the sql statement
		statement, err := db.Prepare(sql)
		_ = CheckErr(err, true)
		_, err = statement.Exec()
		_ = CheckErr(err, true)
		id++
		statement.Close()
	}
	readFile.Close()
}
