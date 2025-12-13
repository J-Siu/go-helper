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

var (
	StrAny     = logger.StrAny
	TimeFormat = "2006-01-02 15:04:05"
	logger     = New()
)

func New() *EzLog   { return new(EzLog).New() }
func Clear() *EzLog { return logger.Clear() }

// Enable/Disable json indent on `data`
func EnableJsonIndent(enable bool) *EzLog { return logger.EnableJsonIndent(enable) }

// Enable/Disable trim on `data`
func EnableTime(enable bool) *EzLog { return logger.EnableTime(enable) }

// Enable/Disable trim on message
func EnableTrim(enable bool) *EzLog { return logger.EnableTrim(enable) }

// Get log level
func GetLogLevel() EzLogLevel { return logger.GetLogLevel() }

// Get log level prefix enable or not
func GetLogLevelPrefix() bool { return logger.GetLogLevelPrefix() }

// Enable/Disable skip empty
func GetSkipEmpty() bool { return logger.GetSkipEmpty() }

// Set all log func to use fmt.Println()
func SetDefaultOutFun() *EzLog { return logger.SetDefaultOutFun() }

// Set DateTime function
func SetDateTimeFunc(f FuncDateTime) *EzLog { return logger.SetDateTimeFunc(f) }

// Set default DateTime function
func SetDefaultDateTimeFunc() *EzLog { return logger.SetDefaultDateTimeFunc() }

// Set log level
func SetLogLevel(level EzLogLevel) *EzLog { return logger.SetLogLevel(level) }

// Enable/Disable log level prefix
func SetLogLevelPrefix(enable bool) *EzLog { return logger.SetLogLevelPrefix(enable) }

// Set Name Postfix
func SetNamePostfixChar(char rune) *EzLog { return logger.SetNamePostfixChar(char) }

// Set Name Postfix to default (':')
func SetNamePostfixCharDefault() *EzLog { return logger.SetNamePostfixCharDefault() }

// Set out function
func SetOutFunc(f FuncOut) *EzLog { return logger.SetOutFunc(f) }

// Enable/Disable skip empty
func SetSkipEmpty(enable bool) *EzLog { return logger.SetSkipEmpty(enable) }

// -- Set log message level

// Log message as level
func LogL(level EzLogLevel) *EzLog { return logger.LogL(level) }

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

// --- Output

func Out() *EzLog      { return logger.Out() }
func String() string   { return logger.String() }
func StringP() *string { return logger.StringP() }

// --- Build log message

// enable/disable message log level prefix.
func Lp(enable bool) *EzLog { return logger.Lp(enable) }

// enable/disable message log level prefix. (Alias for Lp())
func LogPrefix(enable bool) *EzLog { return logger.Lp(enable) }

// Skip current message if `Msg` is empty. Current msg only. (shorthand of MsgSkipEmpty())
func Se() *EzLog { return logger.Se() }

// Skip current message if `Msg` is empty. Current msg only.
func SkipEmpty() *EzLog { return logger.Se() }

// --- base logging functions

// Append character/rune to message
func C(ch rune) *EzLog { return logger.C(ch) }

// Add newline to message.
func L() *EzLog { return logger.L() }

// Add msg to log
func M(date any) *EzLog { return logger.M(date) }

// Add : after data
func N(data any) *EzLog { return logger.N(data) }

// --- Shorthand

// Add new line before data (shorthand for L().M())
func Lm(data any) *EzLog { return logger.L().M(data) }

// Add new line after data (shorthand for M().L())
func Ml(data any) *EzLog { return logger.M(data).L() }

// Add new line before data and : after (shorthand for L().N())
func Ln(data any) *EzLog { return logger.L().N(data) }

// Add : and new line after data (shorthand for N().L())
func Nl(data any) *EzLog { return logger.N(data).L() }

// --- Expressive func name

// Add new line to log (alias of L())
func NewLine() *EzLog { return logger.L() }

// Add msg to log (alias of M())
func Msg(data any) *EzLog { return logger.M(data) }

// Add : after data (alias of N())
func Name(data any) *EzLog { return logger.N(data) }

// --- Expressive alias

// Add new line before data (alias of L().M())
func NewLineMsg(data any) *EzLog { return logger.Lm(data) }

// Add new line after data (alias of M().L())
func MsgNewLine(data any) *EzLog { return logger.Ml(data) }

// Add new line before data and : after (alias of L().N())
func NewLineName(data any) *EzLog { return logger.Ln(data) }

// Add : and new line after data (alias of N().L())
func NameNewLine(data any) *EzLog { return logger.Nl(data) }

// -- Other shorthand

// Add tab to message.
func T() *EzLog { return logger.T() }

// Add tab to message
func Tab() *EzLog { return logger.T() }

// Add "End" to message. (shorthand for M("End"))
func TxtEnd() *EzLog { return logger.TxtEnd() }

// Add "Start" to message. (shorthand for M("Start"))
func TxtStart() *EzLog { return logger.TxtStart() }

// Add "OK"/"Fail" to message.
func Ok(data bool) *EzLog { return logger.Ok(data) }

// Add "Success"/"Fail" to message.
func Success(data bool) *EzLog { return logger.Success(data) }

// Add "Yes"/"No" to message.
func YesNo(data bool) *EzLog { return logger.YesNo(data) }
