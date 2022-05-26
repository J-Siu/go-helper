package helper

// Type Err is a string with error interface
type Err string

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
