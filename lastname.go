package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getLastname() string {
	var name string
	r := rnd.Intn(lastNameCount)
	sql := "select name from lastname where id = " + strconv.Itoa(r) + ";"
	statement, err := db.Prepare(sql)
	_ = CheckErr(err, true)
	rows, err := statement.Query()
	_ = CheckErr(err, true)

	for rows.Next() {
		rows.Scan(&name)
	}
	if doMale || doFemale {
		fmt.Print(" ")
	}
	rows.Close()
	statement.Close()
	name = strings.ToUpper(name)
	return name
}
