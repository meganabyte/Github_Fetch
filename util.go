package util

import (
	"time"
)

func timeDiffInDays(day string) int {
	today := time.Now()
	layout := "2006-01-02"
	t, _ := time.Parse(layout, day)
	daysDiff := (today.Sub(t).Hours())/24
	return int(daysDiff)
}

func AddToMap(aMap map[int]int, time string) {
	index := 365 - timeDiffInDays(time)
	if index > 0 {
		i, ok := aMap[index]
		if ok == false {
			aMap[index] = 1
		} else {
			aMap[index] = i + 1
		}
	}
}

func ComputeContr(mIssues map[int]int, mPulls map[int]int, mCommits map[int]int) [365]int {
	var result [365]int
	arrIssues := getContr(mIssues)
	arrPulls := getContr(mPulls)
	arrCommits := getContr(mCommits)
	for i := 0; i < 365; i++ {
		result[i] = arrIssues[i] + arrPulls[i] + arrCommits[i]
	}
	return result
}

func getContr(m map[int]int) [365]int {
	var arr [365]int
	for key, value := range m {
		arr[key] = value
	}
	return arr
}
