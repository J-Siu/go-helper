package helper

//
type MyArray[T any] []T

// Return true if ErrsType array is empty
func (self *MyArray[T]) Empty() bool {
	return len(*self) == 0
}

// Return true if ErrsType array is not empty
func (self *MyArray[T]) NotEmpty() bool {
	return len(*self) > 0
}

// Clear the ErrsType array
func (self *MyArray[T]) Clear() {
	self = &MyArray[T]{}
}

func (self *MyArray[T]) Add(t T) *MyArray[T] {
	*self = append(*self, t)
	return self
}
