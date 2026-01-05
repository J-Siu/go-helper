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

type StateFunc[T any] func() *State[T]

type State[T any] struct {
	basestruct.Base

	Data T
	Name string // State name

	Next StateFunc[T]
	Post StateFunc[T] // if not nil, always run after Next
	Pre  StateFunc[T] // if not nil, always run before Next

	OnErr         StateFunc[T]
	OnErrContinue bool
}

func (t *State[T]) Run(f StateFunc[T]) *State[T] {
	prefix := t.MyType + ".Loop"
	ezlog.Debug().N(prefix).TxtStart().Out()
	// state machine loop
	t.Next = f
	for t.Next != nil {
		t.Err = nil
		t.
			runState(t.Pre).
			runState(t.Next).
			runState(t.Post)
		if t.Err != nil && !t.OnErrContinue {
			break
		}
	}
	ezlog.Debug().N(prefix).TxtEnd().Out()
	return t
}

func (t *State[T]) runState(f StateFunc[T]) *State[T] {
	if t.Err == nil || t.OnErrContinue {
		if f != nil {
			ezlog.Debug().N(t.MyType).N(t.Name).TxtStart().Out()
			f()
			ezlog.Debug().N(t.MyType).N(t.Name).TxtEnd().Out()
		}
		if t.Err != nil && t.OnErr != nil {
			t.OnErr()
		}
	}
	return t
}
