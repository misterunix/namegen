package main

type firstname struct {
	ID   int
	Name string
}

type lastname struct {
	ID   int
	Name string
}

type lastnameFreq struct {
	ID       int
	Name     string
	Rank     int
	Count    int
	White    float64
	Black    float64
	Hispanic float64
	Asian    float64
	Indian   float64
}

type firstnameFreq struct {
	ID         int
	Name       string
	Count      int
	Percentage float64
	AsString   string
}

type ModifiedNameFreq struct {
	ID             int
	Name           string
	PercentageHigh float64
	PercentageLow  float64
}
