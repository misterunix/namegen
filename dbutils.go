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

// Check if a table exists. Sets createnewdb to true if it does not
func checkIfTableExists(t string) {
	sql := "SELECT name FROM sqlite_master WHERE type='table' AND name='" + t + "';"
	statement, err := db.Prepare(sql)
	_ = CheckErr(err, true)
	rows, err := statement.Query()
	_ = CheckErr(err, true)
	var rc int
	var tmpstring string
	for rows.Next() {
		rc++
		rows.Scan(&tmpstring)
	}
	rows.Close()
	statement.Close()
	if rc == 0 {
		fmt.Println("Table:", t, "does not exist. Creating a new DB with defaults.")
		createnewdb = true
	}

}
