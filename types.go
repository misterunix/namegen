package main

type genericName struct {
	ID   int
	Name string
}

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

type nameFreq struct {
	ID         int
	Name       string
	Count      int
	Percentage float64
	AsString   string
}

type lastNameSimpleFreq struct {
	ID         int
	Name       string
	Count      int
	Percentage float64
}

type lastnamefreq struct {
	ID           int
	Name         string  // The surname
	Rank         int     // OVerall rank
	Count        int     // Number of occurrences
	Prop100k     float64 // Proportion per 100,000
	Cum_prop100k float64 // Cumulative proportion per 100,000
	Pctwhite     float64 // Percent non-Hispanic white
	Pctblack     float64 // Percent non-Hispanic black
	Pctapi       float64 // Percent non-Hispanic Asian/Pacific Islander
	Pctaian      float64 // Percent non-Hispanic American Indian/Alaska Native
	Pct2prace    float64 // Percent dual race
	Pcthispanic  float64 // Percent Hispanic
}

type lastnamefreq2 struct {
	ID   int
	Name string // The surname
	PH   string
	PL   string
}

type ModifiedNameFreq struct {
	ID             int
	Name           string
	PercentageHigh float64
	PercentageLow  float64
}
