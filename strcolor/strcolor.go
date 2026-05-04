/*
The MIT License

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

// Add color to string
package strcolor

const (
	StrColorReset = "\033[0m"

	StrColorBlue    = "\033[34m"
	StrColorCyan    = "\033[36m"
	StrColorGray    = "\033[37m"
	StrColorGreen   = "\033[32m"
	StrColorMagenta = "\033[35m"
	StrColorRed     = "\033[31m"
	StrColorWhite   = "\033[97m"
	StrColorYellow  = "\033[33m"
)

func Blue(s string) string    { return StrColorBlue + s + StrColorReset }
func Cyan(s string) string    { return StrColorCyan + s + StrColorReset }
func Gray(s string) string    { return StrColorGray + s + StrColorReset }
func Green(s string) string   { return StrColorGreen + s + StrColorReset }
func Magenta(s string) string { return StrColorMagenta + s + StrColorReset }
func Red(s string) string     { return StrColorRed + s + StrColorReset }
func White(s string) string   { return StrColorWhite + s + StrColorReset }
func Yellow(s string) string  { return StrColorYellow + s + StrColorReset }
