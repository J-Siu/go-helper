/*
Copyright © 2026 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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

// State machine
package state

import (
	"github.com/J-Siu/go-helper/v2/basestruct"
	"github.com/J-Siu/go-helper/v2/ezlog"
)

type StateFunc[T any] func(state *State[T]) *State[T]

type State[T any] struct {
	basestruct.Base
	Data  T
	Name  string // State name
	Next  StateFunc[T]
	OnErr StateFunc[T]
	Post  StateFunc[T]
	Pre   StateFunc[T]
}

func (t *State[T]) Run(f StateFunc[T], onErrBreak bool) *State[T] {
	prefix := t.MyType + ".Loop"
	ezlog.Debug().N(prefix).TxtStart().Out()
	// state machine loop
	t.Next = f
	for t.Next != nil {
		t.Err = nil
		if t.Pre != nil {
			t.Pre(t)
		}
		t.Next(t)
		if t.Post != nil {
			t.Post(t)
		}
		ezlog.Info().N(t.Name).M("Done").Out()
		if t.Err != nil {
			if t.OnErr != nil {
				t.OnErr(t)
			}
			if onErrBreak {
				break
			}
		}
	}
	ezlog.Debug().N(prefix).TxtEnd().Out()
	return t
}
