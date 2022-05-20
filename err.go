package helper

// Err is a string with error interface
type Err string

// ErrArrayT is array of errors
type ErrsT []error

// `error` interface`
func (self Err) Error() string {
	return string(self)
}

// `stringer` interface
func (self *Err) String() string {
	return string(*self)
}

func (self *Err) StringP() *string {
	return (*string)(self)
}

func (self *ErrsT) NotEmpty() bool {
	return len(*self) > 0
}
