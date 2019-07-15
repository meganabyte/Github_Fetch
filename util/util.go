package util

import (
	"time"
	"math"
)

type Util interface {
	timeDiffInDays(day string)
	AddToMap(aMap map[int]int, time string)
	ComputeContr(mIssues map[int]int, mPulls map[int]int, mCommits map[int]int)
	getContr(m map[int]int)
}

const numDays = 366
type contrArray = [numDays]int

// returns the difference in days between today and the given day
func timeDiffInDays(day string) int {
	today := time.Now()
	layout := "2006-01-02T15:04:05Z07:00" 
	t, _ := time.Parse(layout, day)
	daysDiff := (today.Sub(t).Hours()) / 24
	return int(math.Round(daysDiff))
}

// given a map and time, computes the index and adds mapping between an index and count
// at that index
func AddToMap(aMap map[int]int, time string) {
	index := 365 - timeDiffInDays(time)
	if index > 0 {
		i, ok := aMap[index]
		if !ok {
			aMap[index] = 1
		} else {
			aMap[index] = i + 1
		}
	}
}

// given issues, pulls, and commits, returns a record of all user contributions over the past year
func ComputeContr(mIssues map[int]int, mPulls map[int]int, mCommits map[int]int) contrArray {
	var result contrArray
	arrIssues := getContr(mIssues)
	arrPulls := getContr(mPulls)
	arrCommits := getContr(mCommits)
	for i := 0; i < numDays; i++ {
		result[i] = arrIssues[i] + arrPulls[i] + arrCommits[i]
	}
	return result
}

// given a map, returns a record of a subset of user contributions over the past year
func getContr(m map[int]int) contrArray {
	var arr contrArray
	for key, value := range m {
		arr[key] = value
	}
	return arr
}
