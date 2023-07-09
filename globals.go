package main

import (
	"database/sql"
	"math/rand"
)

var db *sql.DB

var rnd *rand.Rand

var femaleNamesFreq []ModifiedNameFreq
var maleNamesFreq []ModifiedNameFreq

var createnewdb bool
var doMale bool
var doFemale bool
var doLastName bool
var doPercent bool
