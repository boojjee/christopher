package helpers

import ()

func ConvertPoint(distance float64, constant_point float64) float64 {
	return (distance / 1000) * constant_point
}
