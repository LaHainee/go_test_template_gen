package pointer

func To[T any](value T) *T {
	return &value
}

func Val[T comparable](target *T) T {
	var result T
	if target == nil {
		return result
	}

	return *target
}
