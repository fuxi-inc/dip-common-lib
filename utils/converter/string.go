package converter

import "strconv"

type StringConverter interface {
	ToString() string
}

func StringToInt64(s string) int64 {
	if s == "" {
		return 0
	}
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func StringToInt32(s string) int32 {
	if s == "" {
		return int32(0)
	}
	i, _ := strconv.Atoi(s)
	return int32(i)
}

func StringToInt16(s string) int16 {
	if s == "" {
		return int16(0)
	}
	i, _ := strconv.Atoi(s)
	return int16(i)
}
func StringToInt8(s string) int8 {
	if s == "" {
		return int8(0)
	}
	i, _ := strconv.Atoi(s)
	return int8(i)
}

func StringToBool(s string) bool {
	if s == "" {
		return false
	}
	i, _ := strconv.ParseBool(s)
	return i
}

func StringToInt(s string) int {
	if s == "" {
		return 0
	}
	i, _ := strconv.Atoi(s)
	return i
}

func StringToFloat64(s string) float64 {
	if s == "" {
		return float64(0)
	}
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func StringToFloat32(s string) float32 {
	if s == "" {
		return float32(0)
	}
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

func StringInSlice(s string, slice []string) bool {
	if len(slice) == 0 || s == "" {
		return false
	}
	for _, v := range slice {
		if s == v {
			return true
		}
	}
	return false
}
