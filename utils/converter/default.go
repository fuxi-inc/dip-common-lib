package converter

func GetOrDefault(res interface{}, defaultValue interface{}) interface{} {
	if res == nil {
		return defaultValue
	}
	return &res
}

func GetOrDefaultString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func GetOrDefaultInt32(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func GetOrDefaultInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func GetOrDefaultFloat64(i *float64) float64 {
	if i == nil {
		return 0
	}
	return *i
}
