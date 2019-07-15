package util_test 

import (
	"time"
	"testing"
	"util"
	"fmt"
)

func TestTimeDiffInDays(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02T15:04:05Z07:00")
	diff := util.TimeDiffInDays(yesterday)
	if diff != 1 {
		fmt.Printf("Error in TimeDiffInDays: got difference of %v, want %v", diff, 1)
	}
}

func TestAddToMap(t *testing.T) {
	aMap := make(map[int]int)
	time := time.Now().AddDate(-1, 0, -1).Format("2006-01-02T15:04:05Z07:00")
	util.AddToMap(aMap, time)
	if len(aMap) != 0 {
		fmt.Printf("Error in AddToMap: added mapping for contribution made more than a year ago")
	}
}

func TestComputeContr(t *testing.T) {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	m3 := make(map[int]int)
	m1[0] = 1
	m2[0] = 1
	m3[0] = 1
	arr := util.ComputeContr(m1, m2, m3)
	if arr[0] != 3 {
		fmt.Printf("Error in ComputeContr")
	}
}

func TestGetContr(t *testing.T) {
	m := make(map[int]int)
	m[0] = 1
	arr := util.GetContr(m)
	if arr[0] != 1 {
		fmt.Printf("Error in GetContr")
	}
}
