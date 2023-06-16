package converter

import "strconv"

type Float32Converter interface {
	ToFloat32() float32
}

type Float64Converter interface {
	ToFloat64() float64
}

func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}
