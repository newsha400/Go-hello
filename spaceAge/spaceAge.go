package main

import (
	"fmt"
)

//number of seconds in a sloar year
var secondsInYear = 31536000
var ageInSec = 630720000
var ageInYears float32

func findAgeInYears(ageInSec int, secondsInAYear int) float32 {
	return float32(ageInSec / secondsInYear)
}

func solarAge() {
	ageInYears = findAgeInYears(ageInSec, secondsInYear)
	fmt.Printf("%f", ageInYears)
}
