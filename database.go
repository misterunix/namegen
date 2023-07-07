package main

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"

	_ "github.com/glebarez/go-sqlite"
)

func UpdateRow(table string, v map[string]any, where string) string {

	var sql1, sql2, sql3 string
	sql1 = "UPDATE " + table + " SET "

	for p, c := range v {
		switch c.(type) {
		case int:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case int8:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case int16:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case int32:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case int64:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint8:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint16:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint32:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case uint64:
			sql2 += p + " = " + fmt.Sprintf("%d", c) + ","
		case string:
			sql2 += p + " = " + "'" + fmt.Sprintf("%s", c) + "'" + ","
		case float32:
			sql2 += p + " = " + fmt.Sprintf("%f", c) + ","
		case float64:
			sql2 += p + " = " + fmt.Sprintf("%f", c) + ","
		case bool:
			sql2 += p + " = " + fmt.Sprintf("%v", c) + ","
		default:
			return ""
		}

	}

	sql3 = sql1 + sql2 + " WHERE " + where + ";"
	return sql3

}

func RemoveRow(table string, where string) string {

	var sql1, sql2 string
	sql1 = "DELETE FROM " + table + " WHERE "
	sql2 = sql1 + where + ";"
	return sql2

}

func InsertIntoTable(table string, s interface{}) string {

	var middlesql1 string
	var middlesql2 string

	var reflectedValue reflect.Value = reflect.ValueOf(s)

	middlesql1 = "INSERT INTO " + table + " ("
	middlesql2 = ")VALUES("
	for i := 0; i < reflectedValue.NumField(); i++ {

		varName := reflectedValue.Type().Field(i).Name
		varType := reflectedValue.Type().Field(i).Type
		varValue := reflectedValue.Field(i).Interface()

		middlesql1 += varName + ","

		// This is my normal way of working with reflect. Strings may be slower but easier to read.
		switch varType.Kind() {
		case reflect.Int:
			middlesql2 += fmt.Sprintf("%d", varValue.(int)) + ","
		case reflect.Int8:
			middlesql2 += fmt.Sprintf("%d", varValue.(int8)) + ","
		case reflect.Int16:
			middlesql2 += fmt.Sprintf("%d", varValue.(int16)) + ","
		case reflect.Int32:
			middlesql2 += fmt.Sprintf("%d", varValue.(int32)) + ","
		case reflect.Int64:
			middlesql2 += fmt.Sprintf("%d", varValue.(int64)) + ","
		case reflect.Uint:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint)) + ","
		case reflect.Uint8:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint8)) + ","
		case reflect.Uint16:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint16)) + ","
		case reflect.Uint32:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint32)) + ","
		case reflect.Uint64:
			middlesql2 += fmt.Sprintf("%d", varValue.(uint64)) + ","
		case reflect.String:
			middlesql2 += "'" + varValue.(string) + "',"
		case reflect.Float32:
			middlesql2 += fmt.Sprintf("%f", varValue.(float64)) + ","
		case reflect.Float64:
			middlesql2 += fmt.Sprintf("%f", varValue.(float64)) + ","
		case reflect.Bool:
			middlesql2 += fmt.Sprintf("%v", varValue.(bool)) + ","
		default:
			return ""
		}
	}

	middlesql1 = middlesql1[:len(middlesql1)-1]
	middlesql2 = middlesql2[:len(middlesql2)-1] + ");"
	yyy := middlesql1 + middlesql2
	return yyy
}

// Create table based on struct.
// This called by CreateDBtable
// Retuns the sql statement as a string.
// This is a work in progress.
func CreateTableFromStruct(table string, s interface{}) string {

	var reflectedValue reflect.Value = reflect.ValueOf(s) // reflect the struct (interface)

	var sqlstatement string

	//os.Remove("db/savages.db")

	sqlstatement1 := "CREATE TABLE " + table + " ("
	for i := 0; i < reflectedValue.NumField(); i++ {
		var vt string
		varName := reflectedValue.Type().Field(i).Name // get the name of the field
		sqlstatement += "," + varName + " "
		varType := reflectedValue.Type().Field(i).Type // get the type of the field

		// Did this differnt than the other reflect code. This is a work in progress.
		switch varType.Kind() {
		case reflect.Int:
			if varName == "ID" { // detect if the field is the ID field
				vt = "INTEGER NOT NULL PRIMARY KEY"
			} else {
				vt = "INTEGER"
			}
		case reflect.Int8:
			vt = "INTEGER"
		case reflect.Int16:
			vt = "INTEGER"
		case reflect.Int32:
			vt = "INTEGER"
		case reflect.Int64:
			vt = "INTEGER"
		case reflect.Uint:
			vt = "INTEGER"
		case reflect.Uint8:
			vt = "INTEGER"
		case reflect.Uint16:
			vt = "INTEGER"
		case reflect.Uint32:
			vt = "INTEGER"
		case reflect.Uint64:
			vt = "INTEGER"
		case reflect.String:
			vt = "TEXT"
		case reflect.Float64:
			vt = "REAL"
		case reflect.Float32:
			vt = "REAL"
		case reflect.Bool:
			vt = "INTEGER"
		}
		sqlstatement += vt
	}

	// such a crappy way to do this. Return to this at a later date.
	sqlstatement = sqlstatement[1:] // remove the first comma
	sqlstatement += ");"
	sqlstatement = sqlstatement1 + sqlstatement

	return sqlstatement
}

// Remove a table from the database
// table is the name of the table
func DropTable(d *sql.DB, table string) {
	//fmt.Println("Drop table:", table)
	s := fmt.Sprintf("DROP TABLE IF EXISTS %s;", table)
	//fmt.Println(s)
	statement, _ := d.Prepare(s)
	_, err := statement.Exec()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Create a new DB
// dbtable is the name of the table
// sstruct is the struct that will be used to create the table
func CreateDBtable(d *sql.DB, dbtable string, sstruct interface{}) {

	s := CreateTableFromStruct(dbtable, sstruct)
	//fmt.Println(s)
	statement, err := d.Prepare(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	statement.Close()

}
