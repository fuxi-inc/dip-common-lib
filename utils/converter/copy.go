package converter

func Float32CopyPtr(v float32) *float32 {
	vCpy := v
	return &vCpy
}

func Float64CopyPtr(v float64) *float64 {
	vCpy := v
	return &vCpy
}

func IntCopyPtr(v int) *int {
	vCpy := v
	return &vCpy
}

func Int8CopyPtr(v int8) *int8 {
	vCpy := v
	return &vCpy
}

func Int16CopyPtr(v int16) *int16 {
	vCpy := v
	return &vCpy
}

func Int32CopyPtr(v int32) *int32 {
	vCpy := v
	return &vCpy
}

func Int64CopyPtr(v int64) *int64 {
	vCpy := v
	return &vCpy
}

func Uint32CopyPtr(v uint32) *uint32 {
	vCpy := v
	return &vCpy
}

func Uint64CopyPtr(v uint64) *uint64 {
	vCpy := v
	return &vCpy
}

func StringCopyPtr(v string) *string {
	vCpy := v
	return &vCpy
}

func BoolCopyPtr(v bool) *bool {
	vCpy := v
	return &vCpy
}

func Int64PtrCopyPtr(v *int64) *int64 {
	if v == nil {
		return nil
	} else {
		vCpy := *v
		return &vCpy
	}
}

func PtrCopyFloat32(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrCopyFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrCopyInt(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func PtrCopyInt8(v *int8) int8 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrCopyInt16(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrCopyInt32(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrCopyInt64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrCopyUint32(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrCopyUint64(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

func PtrCopyString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func PtrCopyBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}
