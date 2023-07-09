package main

import (
	"fmt"
	"math"
)

func getPercentName() string {
	var name string

	rn := math.Pow(rnd.Float64(), 0.5) * 100.0
	if doFemale {
		for _, k := range femaleNamesFreq {
			if rn <= k.PercentageHigh && rn >= k.PercentageLow {
				fmt.Println(k.Name)
				name = k.Name
				break
			}
		}
	}

	if doMale {
		for _, k := range maleNamesFreq {
			if rn <= k.PercentageHigh && rn >= k.PercentageLow {
				fmt.Println(k.Name)
				name = k.Name
				break
			}
		}

	}

	return name
}
