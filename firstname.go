package main

import (
	"math"
	"strconv"
	"strings"
)

func getPercentName() string {
	var name string

	//rn := math.Pow(rnd.Float64(), 0.5) * 100.0
	rn := math.Sqrt(rnd.Float64()) * 100.0
	if doFemale {
		for _, k := range femaleNamesFreq {
			if rn <= k.PercentageHigh && rn >= k.PercentageLow {
				name = k.Name
				break
			}
		}
	}

	if doMale {
		for _, k := range maleNamesFreq {
			if rn <= k.PercentageHigh && rn >= k.PercentageLow {
				name = k.Name
				break
			}
		}

	}
	name = strings.ToUpper(name)
	return name
}

func getMaleFirstname() string {
	var name string
	r := rnd.Intn(maleCount)
	sql := "select name from firstnamemale where id = " + strconv.Itoa(r) + ";"
	statement, err := db.Prepare(sql)
	_ = CheckErr(err, true)
	rows, err := statement.Query()
	_ = CheckErr(err, true)

	for rows.Next() {
		rows.Scan(&name)
	}
	rows.Close()
	statement.Close()
	name = strings.ToUpper(name)
	return name
}

func getFemaleFirstname() string {
	var name string
	r := rnd.Intn(femaleCount)
	sql := "select name from firstnamefemale where id = " + strconv.Itoa(r) + ";"
	statement, err := db.Prepare(sql)
	_ = CheckErr(err, true)
	rows, err := statement.Query()
	_ = CheckErr(err, true)
	for rows.Next() {
		rows.Scan(&name)
	}
	rows.Close()
	statement.Close()
	name = strings.ToUpper(name)
	return name
}
