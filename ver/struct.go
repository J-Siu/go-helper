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

package ver

import "fmt"

type Version struct {
	major  int
	minor  int
	patch  int
	prefix string
}

// prefix default to "v"
func (t *Version) New() *Version {
	t.prefix = "v"
	return t
}

// Set Prefix of version string
func (t *Version) Prefix(v string) *Version {
	t.prefix = v
	return t
}

// Return version in string form
func (t *Version) String() string {
	return fmt.Sprintf("%s%d.%d.%d", t.prefix, t.major, t.minor, t.patch)
}

// Set Major
func (t *Version) Major(v int) *Version {
	t.major = v
	return t
}

// Set Minor
func (t *Version) Minor(v int) *Version {
	t.minor = v
	return t
}

// Set Patch
func (t *Version) Patch(v int) *Version {
	t.patch = v
	return t
}

// shorthand for setting Major
func (t *Version) M(v int) *Version {
	t.major = v
	return t
}

// shorthand for setting Minor
func (t *Version) N(v int) *Version {
	t.minor = v
	return t
}

// shorthand for setting Patch
func (t *Version) P(v int) *Version {
	t.patch = v
	return t
}
