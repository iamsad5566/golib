package slice

import "github.com/iamsad5566/golib"

func Contains[T golib.DataType](s []T, target T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == target {
			return true
		}
	}
	return false
}
