package array

type Array[T any] []T

// Return true if ErrsType array is empty
func (a *Array[T]) Empty() bool {
	return len(*a) == 0
}

func (a *Array[T]) Len() int {
	return len(*a)
}

// Return true if ErrsType array is not empty
func (a *Array[T]) NotEmpty() bool {
	return len(*a) > 0
}

// Clear the ErrsType array
func (a *Array[T]) Clear() *Array[T] {
	*a = nil
	return a
}

func (a *Array[T]) Add(t T) *Array[T] {
	*a = append(*a, t)
	return a
}
