package libutils

import (
	"math"
)

func Round(num float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Round(num*shift) / shift
}
