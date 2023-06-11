package xstrconv

import "strconv"

func ParseFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func ParseFloat32(s string) (float32, error) {
	result, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}

	return float32(result), nil
}
