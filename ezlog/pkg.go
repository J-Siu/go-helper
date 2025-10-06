/*
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

// ezlog - A simple log mapping module
//
//	-2. Log
//	-1. Disable
//	0. Emerg
//	1. Alert
//	2. Crit
//	3. Err
//	4. Warning
//	5. Notice
//	6. Info
//	7. Debug
//	8. Trace
package ezlog

var logger = New()

func New() *EzLog { return new(EzLog).New() }

// Get log level
func GetLogLevel() Level { return logger.GetLogLevel() }

// Get log level prefix enable or not
func GetLogLevelPrefix() bool { return logger.GetLogLevelPrefix() }

// Set log level
func SetLogLevel(level Level) *EzLog { return logger.SetLogLevel(level) }

// Enable/Disable log level prefix
func SetLogLevelPrefix(enable bool) *EzLog { return logger.SetLogLevelPrefix(enable) }

// Set all log func to use fmt.Print()
func SetOutPrint() *EzLog { return logger.SetOutPrint() }

// Set all log func to use fmt.Println()
func SetOutPrintLn() *EzLog { return logger.SetOutPrintLn() }

// Enable/Disable trim on message
func SetTrim(enable bool) *EzLog { return logger.SetTrim(enable) }

func String() string   { return logger.String() }
func StringP() *string { return logger.StringP() }

// Log message as level
func LogL(level Level) *EzLog { return logger.LogL(level) }

// Log message without log level
func Log() *EzLog { return logger.Log() }

// Log message as `EMERG`
func Emerg() *EzLog { return logger.Emerg() }

// Log message as `ALERT`
func Alert() *EzLog { return logger.Alert() }

// Log message as `CRIT`
func Crit() *EzLog { return logger.Crit() }

// Log message as `ERR`
func Err() *EzLog { return logger.Err() }

// Log message as `WARNING`
func Warning() *EzLog { return logger.Warning() }

// Log message as `NOTICE`
func Notice() *EzLog { return logger.Notice() }

// Log message as `INFO`
func Info() *EzLog { return logger.Info() }

// Log message as `DEBUG`
func Debug() *EzLog { return logger.Debug() }

// Log message as `TRACE`
func Trace() *EzLog { return logger.Trace() }
