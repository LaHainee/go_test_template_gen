package set

type Set[T comparable] map[T]struct{}

func New[T comparable](vals ...T) Set[T] {
	set := Set[T]{}

	for _, val := range vals {
		set.Add(val)
	}

	return set
}

func (set Set[T]) Size() int {
	return len(set)
}

func (set Set[T]) Remove(val T) {
	delete(set, val)
}

func (set Set[T]) Values() []T {
	values := make([]T, 0, len(set))

	for val := range set {
		values = append(values, val)
	}

	return values
}

func (set Set[T]) Add(values ...T) {
	for _, val := range values {
		set[val] = struct{}{}
	}
}

func (set Set[T]) Contains(item T) bool {
	_, ok := set[item]
	return ok
}
