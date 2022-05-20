package helper

type Warning string
type Warnings []Warning

// `error` interface
func (self Warning) Error() string {
	return string(self)
}

func (self *Warning) String() string {
	return string(*self)
}

func (self *Warning) StringP() *string {
	return (*string)(self)
}

func (self *Warnings) NotEmpty() bool {
	return len(*self) > 0
}
