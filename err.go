package helper

// Type Err is a string with error interface
type Err string

// Type ErrT is array of errors
type ErrsT []error

// `error` interface
func (self Err) Error() string {
	return string(self)
}

// Return content of Err as string
func (self *Err) String() string {
	return string(*self)
}

// Return content of Err as *string
func (self *Err) StringP() *string {
	return (*string)(self)
}

// Return true if ErrsT array is empty
func (self *ErrsT) Empty() bool {
	return len(*self) == 0
}

// Return true if ErrsT array is not empty
func (self *ErrsT) NotEmpty() bool {
	return len(*self) > 0
}

// Clear the ErrsT array
func (self *ErrsT) Clear() {
	self = &ErrsT{}
}
