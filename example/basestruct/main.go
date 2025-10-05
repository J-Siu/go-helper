/*
The MIT License

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"

	bigObj "github.com/J-Siu/go-helper/v2/example/basestruct/bigobj"
)

func main() {

	var (
		str  string
		obj1 bigObj.BigObject
	)

	fmt.Println("-- obj1: Using [bigObj.BigObject] with no initialization:")
	str = obj1.SetName("Cube").SetHeight(10).SetLength(10).SetWidth(10).String()
	if obj1.Err != nil {
		fmt.Println(obj1.Err.Error()) // Normal to reach this line
	} else {
		fmt.Println(str)
	}

	fmt.Println("\n-- obj2: Name not set:")
	obj2 := bigObj.New()
	str = obj2.SetHeight(10).SetLength(10).SetWidth(10).String()
	if obj2.Err != nil {
		fmt.Println(obj2.Err.Error()) // Normal to reach this line
	} else {
		fmt.Println(str)
	}

	fmt.Println("\n--obj3: No error:")
	obj3 := bigObj.New()
	str = obj3.SetName("Cube").SetHeight(10).SetLength(10).SetWidth(10).String()
	if obj3.Err != nil {
		fmt.Println(obj3.Err.Error()) // Should not reach this line
	} else {
		fmt.Println(str)
	}

}
