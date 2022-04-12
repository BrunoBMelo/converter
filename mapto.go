package converter

import "fmt"

func To[T any](in interface{}) (T, error) {
	if obj, ok := in.(T); !ok {
		var result T
		return result, fmt.Errorf("cannot tranform interface %T in %T ", in, result)
	} else {
		return obj, nil
	}
}
