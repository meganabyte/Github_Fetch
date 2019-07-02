package util

import (
	"fmt"
	"os"
	"time"
)

// returns the difference in days between today and the given day
func timeDiffInDays(day string) int {
	today := time.Now()
	layout := "2006-01-02"
	t, _ := time.Parse(layout, day)
	daysDiff := (today.Sub(t).Hours()) / 24
	return int(daysDiff)
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
func ComputeContr(mIssues map[int]int, mPulls map[int]int, mCommits map[int]int) [366]int {
	var result [366]int
	arrIssues := getContr(mIssues)
	arrPulls := getContr(mPulls)
	arrCommits := getContr(mCommits)
	for i := 0; i < 366; i++ {
		result[i] = arrIssues[i] + arrPulls[i] + arrCommits[i]
	}
	return result
}

// given a map, returns a record of a subset of user contributions over the past year
func getContr(m map[int]int) [366]int {
	var arr [366]int
	for key, value := range m {
		arr[key] = value
	}
	return arr
}

// given an error, prints error and throws error if error is not nil
func ThrowError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
