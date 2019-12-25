//Package others contains implementation for algorithms
package others

import (
	"strconv"
)

//Sqrt given a number x, find the number z for which z² is most nearly x.
func Sqrt(x float64) float64 {
	z := 1.0
	lastExec := ""

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)

		if lastExec == strconv.FormatFloat(z, 'f', 4, 64) {
			break
		}
		lastExec = strconv.FormatFloat(z, 'f', 4, 64)
	}

	result, _ := strconv.ParseFloat(strconv.FormatFloat(z, 'f', 2, 64), 64)

	return result
}
