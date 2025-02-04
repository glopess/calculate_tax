package helpers

import "math"

func ToTwoPlaces(decimal float64) float64 {
	return math.Round(decimal*100) / 100
}
