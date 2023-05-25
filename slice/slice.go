package slice

func Contains[T int | int8 | int16 | int32 | int64 | float32 | float64](s []T, target T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == target {
			return true
		}
	}
	return false
}
