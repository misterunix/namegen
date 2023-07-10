package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getLastname() string {
	if !doPercent {
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
	} else {
		var name string
		rn := math.Sqrt(rnd.Float64()) * 100.0
		if doPercent {
			for _, k := range surnameFreq {
				if rn <= k.PercentageHigh && rn >= k.PercentageLow {
					name = k.Name
					break
				}
			}
		}
		return name
	}
}

// Load surnames from file
func loadSurnamesFromFile() {
	fmt.Println("\t\tLastnames frequency.")
	fmt.Println("\t\t\tGetting total from surnamefreq.txt")
	readFile, err := os.Open("storage/surnamefreq.txt")
	_ = CheckErr(err, true)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var id int = 0
	var totalCount int = 0
	for fileScanner.Scan() {
		tmp := strings.TrimSpace(fileScanner.Text()) // remove spaces just in case
		sraw := strings.Split(tmp, ",")
		// index 0 is the name
		// index 1 is the rank
		// index 2 is the count
		// index 3 is the proportion per 100,000
		// index 4 is the cumulative proportion per 100,000
		// index 5 is the percent non-Hispanic white
		// index 6 is the percent non-Hispanic black
		// index 7 is the percent non-Hispanic Asian/Pacific Islander
		// index 8 is the percent non-Hispanic American Indian/Alaska Native
		// index 9 is the percent non-Hispanic 2 races
		// index 10 is the percent Hispanic
		ti, err := strconv.Atoi(sraw[2]) // Count, need to place in temp var because off error
		_ = CheckErr(err, true)
		totalCount += ti // Add to total count
		id++
		if id >= surnameLimit {
			break
		}
	}
	readFile.Close()
	fmt.Println("\t\t\tdone.")

	fmt.Println("\t\tCalculating percentages.")

	readFile, err = os.Open("storage/surnamefreq.txt")
	_ = CheckErr(err, true)
	fileScanner = bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lastP float64 = 100.0
	id = 0
	for fileScanner.Scan() {
		tmpln := lastnamefreq2{}
		tmp := strings.TrimSpace(fileScanner.Text()) // remove spaces just in case
		sraw := strings.Split(tmp, ",")
		// index 0 is the name
		// index 1 is the rank
		// index 2 is the count
		// index 3 is the proportion per 100,000
		// index 4 is the cumulative proportion per 100,000
		// index 5 is the percent non-Hispanic white
		// index 6 is the percent non-Hispanic black
		// index 7 is the percent non-Hispanic Asian/Pacific Islander
		// index 8 is the percent non-Hispanic American Indian/Alaska Native
		// index 9 is the percent non-Hispanic 2 races
		// index 10 is the percent Hispanic
		tmpln.ID = id
		tmpln.Name = sraw[0]
		lt, err := strconv.ParseFloat(sraw[2], 64)
		_ = CheckErr(err, true)
		h := lastP
		l := h - (lt/float64(totalCount))*100.0
		tmpln.PH = strconv.FormatFloat(h, 'f', 20, 64)
		tmpln.PL = strconv.FormatFloat(l, 'f', 20, 64)
		lastP = l
		sql := InsertIntoTable("surnames", tmpln)
		_, err = db.Exec(sql)
		_ = CheckErr(err, true)
		id++
		if id >= surnameLimit {
			break
		}
	}
	readFile.Close()
	fmt.Println("\t\t\tdone.")
}

// get the lastname by frequency
func getLastNameFreq() string {
	var name string

	//rn := math.Pow(rnd.Float64(), 0.5) * 100.0
	rn := math.Sqrt(rnd.Float64()) * 100.0
	if doPercent {
		for _, k := range surnameFreq {
			if rn <= k.PercentageHigh && rn >= k.PercentageLow {
				name = k.Name
				break
			}
		}
	}
	return name

}
