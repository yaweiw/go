package funcs

import (
	"fmt"
	"time"

	"github.com/yaweiw/go/structs"
)

// Timeinday is a func
func Timeinday() {
	WEEKDAYS := []structs.WHEN{
		{D: []int{1}, M: "morning", A: "arvo", E: "evening"},
		{D: []int{2}, M: "morning", A: "arvo", E: "evening"},
		{D: []int{3}, M: "morning", A: "arvo", E: "evening"},
		{D: []int{4}, M: "morning", A: "arvo", E: "evening"},
		{D: []int{5}, M: "morning", A: "arvo", E: "evening"},
	}
	fmt.Println(WEEKDAYS)
	WEEKEND := []structs.WHEN{
		{D: []int{6}, M: "morning", A: "arvo", E: "evening"},
		{D: []int{7}, M: "morning", A: "arvo", E: "evening"},
	}
	fmt.Println(WEEKEND)

	var m = map[string]structs.WHEN{
		"MONDAY":    WEEKDAYS[0],
		"TUESDAY":   WEEKDAYS[1],
		"WEDNESDAY": WEEKDAYS[2],
		"THURSDAY":  WEEKDAYS[3],
		"FRIDAY":    WEEKDAYS[4],
	}
	fmt.Println(m)
	w := structs.WHEN{D: []int{29, 30, 31}, M: "morning", A: "arvo", E: "evening"}
	pw := &w
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println(pw.M)
		fmt.Printf("%v", pw.D)
		fmt.Println(pw.D)
	case t.Hour() < 17:
		fmt.Println(pw.A)
		fmt.Printf("%v", pw.D)
		fmt.Println(pw.D)
	default:
		fmt.Println(pw.E)
		fmt.Printf("%v", pw.D)
		fmt.Println(pw.D)
	}
}
