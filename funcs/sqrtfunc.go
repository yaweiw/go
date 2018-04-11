package funcs

import (
	"fmt"
	"math"
)

// Sqrtfunc is a func
func Sqrtfunc(x, y float64) string {
	if x < 0 && x < y {
		return fmt.Sprint(math.Sqrt(-x)) + fmt.Sprint(math.Sqrt(-y)) + "i"
	} // else {
	return fmt.Sprint(math.Sqrt(x))
	//}
}
