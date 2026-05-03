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
	reset = "\033[0m"

	blue    = "\033[34m"
	cyan    = "\033[36m"
	gray    = "\033[37m"
	green   = "\033[32m"
	magenta = "\033[35m"
	red     = "\033[31m"
	white   = "\033[97m"
	yellow  = "\033[33m"
)

func Blue(s *string) string {
	return blue + *s + reset
}
func Cyan(s *string) string {
	return cyan + *s + reset
}
func Gray(s *string) string {
	return gray + *s + reset
}
func Green(s *string) string {
	return green + *s + reset
}
func Magenta(s *string) string {
	return magenta + *s + reset
}
func Red(s *string) string {
	return red + *s + reset
}
func White(s *string) string {
	return white + *s + reset
}
func Yellow(s *string) string {
	return yellow + *s + reset
}
