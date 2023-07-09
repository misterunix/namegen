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

var nameGenCount int // the number of names to generate a name
var maleCount int
var femaleCount int
var lastNameCount int
