package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

func CheckErr(err error, fatal bool) error {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if fatal {
			os.Exit(1)
		}
	}
	return err
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Populate the database with the text from the files.
// This is going to be a messy function and needs to be cleaned up.
func populateDB() {

	//var maleFreqCount int
	//var femaleFreqCount int

	ts := time.Now()
	fmt.Println("\tPopulating the database.")
	fmt.Println("\t\tFemale first names.")

	o := "BEGIN;\n"
	beginstatement, err := db.Prepare(o)
	_ = CheckErr(err, true)
	_, err = beginstatement.Exec()
	_ = CheckErr(err, true)

	readFile, err := os.Open("storage/first-f.txt")
	_ = CheckErr(err, true)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var id int = 0
	for fileScanner.Scan() {
		tmp := strings.TrimSpace(fileScanner.Text())   // remove spaces just in case
		fnf := firstname{id, tmp}                      // create a firstname struct
		sql := InsertIntoTable("firstnamefemale", fnf) // create the sql statement
		statement, err := db.Prepare(sql)
		_ = CheckErr(err, true)
		_, err = statement.Exec()
		_ = CheckErr(err, true)
		id++
	}
	readFile.Close()
	fmt.Println("\t\tdone.")

	fmt.Println("\t\tMale first names.")
	readFile, err = os.Open("storage/first-m.txt")
	_ = CheckErr(err, true)
	fileScanner = bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	id = 0
	for fileScanner.Scan() {
		tmp := strings.TrimSpace(fileScanner.Text()) // remove spaces just in case
		fnm := firstname{id, tmp}                    // create a firstname struct
		sql := InsertIntoTable("firstnamemale", fnm) // create the sql statement
		statement, err := db.Prepare(sql)
		_ = CheckErr(err, true)
		_, err = statement.Exec()
		_ = CheckErr(err, true)
		id++
	}
	readFile.Close()
	fmt.Println("\t\tdone.")

	fmt.Println("\t\tLast names.")
	readFile, err = os.Open("storage/last.txt")
	_ = CheckErr(err, true)
	fileScanner = bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	id = 0
	for fileScanner.Scan() {
		tmp := strings.TrimSpace(fileScanner.Text()) // remove spaces just in case
		lastname := firstname{id, tmp}               // create a firstname struct
		sql := InsertIntoTable("lastname", lastname) // create the sql statement
		statement, err := db.Prepare(sql)
		_ = CheckErr(err, true)
		_, err = statement.Exec()
		_ = CheckErr(err, true)
		id++
	}
	readFile.Close()

	fmt.Println("\t\tMale first names frequency.")
	readFile, err = os.Open("storage/male-new-freq.txt")
	_ = CheckErr(err, true)
	fileScanner = bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	id = 0
	for fileScanner.Scan() {

		tmp := strings.TrimSpace(fileScanner.Text()) // remove spaces just in case
		//split iy on tabs
		sraw := strings.Split(tmp, "\t")
		//fmt.Println(sraw)
		fn := sraw[1]
		//fp, err := strconv.ParseFloat(sraw[2], 64)
		//_ = CheckErr(err, true)
		fp := 0.0
		fcb, err := strconv.ParseInt(sraw[3], 10, 0)
		_ = CheckErr(err, true)
		fc := int(fcb)
		//maleFreqCount += fc
		//var fp float64 = 0.0
		fnm := firstnameFreq{ID: id, Name: fn, Percentage: fp, Count: fc, AsString: sraw[2]} // create a firstname struct
		sql := InsertIntoTable("malefreq", fnm)                                              // create the sql statement
		statement, err := db.Prepare(sql)
		_ = CheckErr(err, true)
		_, err = statement.Exec()
		_ = CheckErr(err, true)
		id++
	}
	readFile.Close()
	fmt.Println("\t\tdone.")

	fmt.Println("\t\tFemale first names frequency.")
	readFile, err = os.Open("storage/female-new-freq.txt")
	_ = CheckErr(err, true)
	fileScanner = bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	id = 0
	for fileScanner.Scan() {

		tmp := strings.TrimSpace(fileScanner.Text()) // remove spaces just in case
		//split iy on tabs
		sraw := strings.Split(tmp, "\t")
		fn := sraw[1]
		//fp, err := strconv.ParseFloat(sraw[2], 64)
		//_ = CheckErr(err, true)
		fp := 0.0
		fcb, err := strconv.ParseInt(sraw[3], 10, 0)
		_ = CheckErr(err, true)
		fc := int(fcb)
		//femaleFreqCount += fc
		//var fp float64 = 0.0
		fnm := firstnameFreq{ID: id, Name: fn, Percentage: fp, Count: fc, AsString: sraw[2]} // create a firstname struct
		sql := InsertIntoTable("femalefreq", fnm)                                            // create the sql statement
		statement, err := db.Prepare(sql)
		_ = CheckErr(err, true)
		_, err = statement.Exec()
		_ = CheckErr(err, true)
		id++
	}
	readFile.Close()
	fmt.Println("\t\tdone.")

	o = "COMMIT;\n"
	beginstatement, err = db.Prepare(o)
	_ = CheckErr(err, true)
	_, err = beginstatement.Exec()
	_ = CheckErr(err, true)

	fmt.Println("\t\tdone.")

	td := time.Since(ts)
	tstr := td.Minutes()
	fmt.Println("\tDone. ", tstr, " minutes.")
}

func checkTableCount(table string) int {
	//fmt.Print("Checking table count: ", table, " : ")
	err := db.Ping()
	_ = CheckErr(err, true)

	count := 0
	sql := "select count(*) from " + table + ";"
	statement, err := db.Prepare(sql)
	e := CheckErr(err, true)
	if e != nil {
		fmt.Println("Error:", e)
	}
	rows, err := statement.Query()
	e = CheckErr(err, true)
	if e != nil {
		fmt.Println("Error:", e)
	}

	for rows.Next() {
		rows.Scan(&count)
	}
	rows.Close()
	//fmt.Print("Count:", c, "\n")
	return count
}

func main() {

	var err error
	var maleCount int
	var femaleCount int
	var lastNameCount int
	var femalFreqCount int
	var maleFreqCount int

	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	flag.BoolVar(&createnewdb, "newdb", false, "Create a new database.")
	flag.BoolVar(&doMale, "m", false, "Generate a male name.")
	flag.BoolVar(&doFemale, "f", false, "Generate a female name.")
	flag.BoolVar(&doLastName, "l", false, "Generate a last name.")
	flag.BoolVar(&doPercent, "p", false, "Use the percentage tables.")
	flag.IntVar(&nameGenCount, "c", 1, "The number of names to generate.")
	flag.Parse()

	// Check if the DB exists. If not, create it.
	if !fileExists("db/names.db") {
		fmt.Println("Database does not exist. Creating a new one.")
		createnewdb = true
	}

	// Open the database. If it doesnt exist, create it as well.
	db, err = sql.Open("sqlite", "db/names.db")
	_ = CheckErr(err, true)
	db.Ping()
	_ = CheckErr(err, true)

	sql := "SELECT name FROM sqlite_master WHERE type='table' AND name='firstnamefemale';"
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
	fmt.Printf("female table: %d\n", rc)
	if rc == 0 {
		createnewdb = true
	}

	sql = "SELECT name FROM sqlite_master WHERE type='table' AND name='firstnamemale';"
	statement, err = db.Prepare(sql)
	_ = CheckErr(err, true)
	rows, err = statement.Query()
	_ = CheckErr(err, true)
	//var rc int
	rc = 0
	for rows.Next() {
		rc++
		rows.Scan(&tmpstring)
	}
	fmt.Printf("male table: %d\n", rc)
	if rc == 0 {
		createnewdb = true
	}

	sql = "SELECT name FROM sqlite_master WHERE type='table' AND name='lastname';"
	statement, err = db.Prepare(sql)
	_ = CheckErr(err, true)
	rows, err = statement.Query()
	_ = CheckErr(err, true)
	//var rc int
	rc = 0
	for rows.Next() {
		rc++
		rows.Scan(&tmpstring)
	}
	fmt.Printf("lastname table: %d\n", rc)
	if rc == 0 {
		createnewdb = true
	}

	sql = "SELECT name FROM sqlite_master WHERE type='table' AND name='femalefreq';"
	statement, err = db.Prepare(sql)
	_ = CheckErr(err, true)
	rows, err = statement.Query()
	_ = CheckErr(err, true)
	//var rc int
	//var tmpstring string
	for rows.Next() {
		rc++
		rows.Scan(&tmpstring)
	}
	fmt.Printf("female frequency table: %d\n", rc)
	if rc == 0 {
		createnewdb = true
	}

	sql = "SELECT name FROM sqlite_master WHERE type='table' AND name='malefreq';"
	statement, err = db.Prepare(sql)
	_ = CheckErr(err, true)
	rows, err = statement.Query()
	_ = CheckErr(err, true)
	//var rc int
	//var tmpstring string
	for rows.Next() {
		rc++
		rows.Scan(&tmpstring)
	}
	fmt.Printf("male frequency table: %d\n", rc)
	if rc == 0 {
		createnewdb = true
	}

	if createnewdb {
		fmt.Println("Creating a new database tables.")
		DropTable(db, "firstnamefemale")
		DropTable(db, "firstnamemale")
		DropTable(db, "femalefreq")
		DropTable(db, "malefreq")
		DropTable(db, "lastname")
		CreateDBtable(db, "firstnamefemale", firstname{})
		CreateDBtable(db, "firstnamemale", firstname{})
		CreateDBtable(db, "lastname", lastname{})
		CreateDBtable(db, "femalefreq", firstnameFreq{})
		CreateDBtable(db, "malefreq", firstnameFreq{})
		populateDB()
		fmt.Println("Done.")
	}

	maleCount = checkTableCount("firstnamemale")
	femaleCount = checkTableCount("firstnamefemale")
	lastNameCount = checkTableCount("lastname")
	femalFreqCount = checkTableCount("femalefreq")
	maleFreqCount = checkTableCount("malefreq")

	fmt.Printf("femaleCount: %d\n", femaleCount)
	fmt.Printf("maleCount: %d\n", maleCount)
	fmt.Printf("lastNameCount: %d\n\n", lastNameCount)
	fmt.Printf("femaleFreqCount: %d\n", femalFreqCount)
	fmt.Printf("maleFreqCount: %d\n\n", maleFreqCount)

	// Load db into memory
	fmt.Println("Loading database into memory.")

	fm := firstnameFreq{}
	sqlfemale := "select * from femalefreq;"
	statement, err = db.Prepare(sqlfemale)
	_ = CheckErr(err, true)
	rows, err = statement.Query()
	_ = CheckErr(err, true)
	var lastP float64 = 100.0
	for rows.Next() {

		rows.Scan(&fm.ID, &fm.Name, &fm.Count, &fm.Percentage, &fm.AsString)
		fm.Percentage, err = strconv.ParseFloat(fm.AsString, 64)
		_ = CheckErr(err, true)
		fm.Percentage *= 100.0
		fm.Percentage = lastP - fm.Percentage

		mfn := ModifiedNameFreq{ID: fm.ID, Name: fm.Name, PercentageHigh: lastP, PercentageLow: fm.Percentage}
		femaleNamesFreq = append(femaleNamesFreq, mfn)
		lastP = fm.Percentage
		if fm.Percentage < 0.0 {
			fm.Percentage = 0.0
		}
	}
	rows.Close()
	statement.Close()

	/*
		fmt.Println("femaleNamesFreq")
		for j, k := range femaleNamesFreq {
			fmt.Println(j, k)
		}
	*/

	fmm := firstnameFreq{}
	sqlmale := "select * from malefreq;"
	statement, err = db.Prepare(sqlmale)
	_ = CheckErr(err, true)
	rows, err = statement.Query()
	_ = CheckErr(err, true)
	lastP = 100.0
	for rows.Next() {

		rows.Scan(&fmm.ID, &fmm.Name, &fmm.Count, &fmm.Percentage, &fmm.AsString)
		//fmt.Println(fmm)
		fmm.Percentage, err = strconv.ParseFloat(fmm.AsString, 64)
		_ = CheckErr(err, true)
		fmm.Percentage *= 100.0
		fmm.Percentage = lastP - fmm.Percentage

		mfn := ModifiedNameFreq{ID: fmm.ID, Name: fmm.Name, PercentageHigh: lastP, PercentageLow: fmm.Percentage}
		maleNamesFreq = append(maleNamesFreq, mfn)

		lastP = fmm.Percentage
		if fm.Percentage < 0.0 {
			fm.Percentage = 0.0
		}

	}
	rows.Close()
	statement.Close()

	/*
		fmt.Println("maleNamesFreq")
		for j, k := range maleNamesFreq {
			fmt.Println(j, k)
		}
	*/
	fmt.Println("Done.")

	for i := 0; i < nameGenCount; i++ {

		if doPercent {
			getPercentName()
		}

		if !doPercent {

			if doMale {
				r := rnd.Intn(maleCount)
				sql := "select name from firstnamemale where id = " + strconv.Itoa(r) + ";"
				statement, err := db.Prepare(sql)
				_ = CheckErr(err, true)
				rows, err := statement.Query()
				_ = CheckErr(err, true)
				var name string
				for rows.Next() {
					rows.Scan(&name)
				}
				fmt.Print(name)
				rows.Close()
			}

			if doFemale {
				r := rnd.Intn(femaleCount)
				sql := "select name from firstnamefemale where id = " + strconv.Itoa(r) + ";"
				statement, err := db.Prepare(sql)
				_ = CheckErr(err, true)
				rows, err := statement.Query()
				_ = CheckErr(err, true)
				var name string
				for rows.Next() {
					rows.Scan(&name)
				}
				fmt.Print(name)
				rows.Close()
			}

			if doLastName {
				r := rnd.Intn(lastNameCount)
				sql := "select name from lastname where id = " + strconv.Itoa(r) + ";"
				statement, err := db.Prepare(sql)
				_ = CheckErr(err, true)
				rows, err := statement.Query()
				_ = CheckErr(err, true)
				var name string
				for rows.Next() {
					rows.Scan(&name)
				}
				if doMale || doFemale {
					fmt.Print(" ")
				}
				fmt.Println(name)
				rows.Close()
			} else {
				fmt.Println()
			}
		}
	}

	defer db.Close()
}
