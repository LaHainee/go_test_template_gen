package slice

func Remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func Reverse[T any](slice []T) []T {
	reversed := make([]T, 0, len(slice))

	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}

	return reversed
}

func Insert[T any](slice []T, index int, value T) []T {
	if len(slice) == index {
		return append(slice, value)
	}

	slice = append(slice[:index+1], slice[index:]...) // index < len(a)
	slice[index] = value

	return slice
}
