package converter

type BoolConverter interface {
	ToBool() bool
}

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
