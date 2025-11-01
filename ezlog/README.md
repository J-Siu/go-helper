# ezlog

A simple log module with Linux log level:

- -2 LOG EzLogLevel = iota - 2 // `LOG` is not exactly a log level. It is for logging regardless of log level
- -1 DISABLED
0. EMERG
1. ALERT
2. CRIT
3. ERR
4. WARNING
5. NOTICE
6. INFO
7. DEBUG
8. TRACE

## Installation

```sh
go get github.com/J-Siu/go-helper/v2
```

## Usage

```go
import "github.com/J-Siu/go-helper/v2/ezlog"
```

## Types and Functions

### Log Level

```go
type Level int8

// log level
const (
  LOG EzLogLevel = iota - 2 // `LOG` is not exactly a log level. It is for logging regardless of log level
  DISABLED
  EMERG
  ALERT
  CRIT
  ERR
  WARNING
  NOTICE
  INFO
  DEBUG
  TRACE
)
```

### EzLog Functions

```go
type OutFunc func(msg *string)

type EzLog struct {
  StrAny         *strany.StrAny `json:"StrAny"`
  logLevel       EzLogLevel
  logLevelPrefix bool
  outFunc        OutFunc
  skipEmpty      bool
  strBuf         []string
  trim           bool
  // msg level
  msgLogLevel       EzLogLevel
  msgNotEmpty       bool
  msgSkipEmpty      bool
  msgLogLevelPrefix bool
}
```

```go
func (t *EzLog) New() *EzLog
// Clear message
func (t *EzLog) Clear() *EzLog
// Get log level
func (t *EzLog) GetLogLevel() EzLogLevel
// Get log level prefix enable or not
func (t *EzLog) GetLogLevelPrefix() bool
// Set log level
func (t *EzLog) SetLogLevel(level EzLogLevel) *EzLog
// Set log level prefix true/false
func (t *EzLog) SetLogLevelPrefix(enable bool) *EzLog
// Set out function
func (t *EzLog) SetOutFunc(f OutFunc) *EzLog
// Set out function to fmt.Print()
func (t *EzLog) SetOutPrint() *EzLog
// Set out function to fmt.Println()
func (t *EzLog) SetOutPrintLn() *EzLog
// Enable/Disable trim on `data`
func (t *EzLog) SetTrim(enable bool) *EzLog
// Enable/Disable trim on `data`
func (t *EzLog) SetSkipEmpty(enable bool) *EzLog
// -- Set log message level
// Log message as `level`
func (t *EzLog) LogL(level EzLogLevel) *EzLog
// Log message without level (no level prefix)
func (t *EzLog) Log() *EzLog
// Log message as `EMERG`
func (t *EzLog) Emerg() *EzLog
// Log message as `ALERT`
func (t *EzLog) Alert() *EzLog
// Log message as `CRIT`
func (t *EzLog) Crit() *EzLog
// Log message as `ERR`
func (t *EzLog) Err() *EzLog
// Log message as `WARN`
func (t *EzLog) Warning() *EzLog
// Log message as `NOTICE`
func (t *EzLog) Notice() *EzLog
// Log message as `INFO`
func (t *EzLog) Info() *EzLog
// Log message as `DEBUG`
func (t *EzLog) Debug() *EzLog
// Log message as `TRACE`
func (t *EzLog) Trace() *EzLog
// --- Output
func (t *EzLog) Dump() *EzLog
func (t *EzLog) Out() *EzLog
func (t *EzLog) String() string
func (t *EzLog) StringP() *string
// --- Build log message
// --- Msg control
// enable/disable message log level prefix.
func (t *EzLog) Lp(enable bool) *EzLog
// skip current message if `Msg` is empty. Current msg only.
func (t *EzLog) Se() *EzLog
// enable/disable trim data. Default to false. Reset each time Out() is called
func (t *EzLog) Tr(enable bool) *EzLog
// enable/disable message log level prefix. (Alias for Lp())
func (t *EzLog) LogPrefix(enable bool) *EzLog
// skip current message if `Msg` is empty. Current msg only. (alias of Se())
func (t *EzLog) SkipEmpty() *EzLog
// enable/disable trim data. Default to false. Reset each time Out() is called. (alias of Trim())
func (t *EzLog) TrimData(enable bool) *EzLog
// --- base logging functions
// Append character/rune to message (shorthand for Sp())
func (t *EzLog) C(ch rune) *EzLog
// Add newline to message.
func (t *EzLog) L() *EzLog { return t.C('
') }
// Add msg to log
func (t *EzLog) M(data any) *EzLog
// Add : after data
func (t *EzLog) N(data any) *EzLog
// --- Shorthand
// Add new line before data (shorthand for L().M())
func (t *EzLog) Lm(data any) *EzLog
// Add new line after data (shorthand for M().L())
func (t *EzLog) Ml(data any) *EzLog
// Add new line before data and : after (shorthand for L().N())
func (t *EzLog) Ln(data any) *EzLog
// Add : and new line after data (shorthand for N().L())
func (t *EzLog) Nl(data any) *EzLog
// --- Expressive func name
// Add new line to log (alias of L())
func (t *EzLog) NewLine() *EzLog
// Add msg to log (alias of M())
func (t *EzLog) Msg(data any) *EzLog
// Add : after data (alias of N())
func (t *EzLog) Name(data any) *EzLog
// --- Expressive alias
// Add new line before data (alias of L().M())
func (t *EzLog) NewLineMsg(data any) *EzLog
// Add new line after data (alias of M().L())
func (t *EzLog) MsgNewLine(data any) *EzLog
// Add new line before data and : after (alias of L().N())
func (t *EzLog) NewLineName(data any) *EzLog
// Add : and new line after data (alias of N().L())
func (t *EzLog) NameNewLine(data any) *EzLog
// --- Other shorthand
// Add tab to message.
func (t *EzLog) T() *EzLog
// Add tab to message. (alias for T())
func (t *EzLog) Tab() *EzLog
// Add "End" to message. (shorthand for M("End"))
func (t *EzLog) TxtEnd() *EzLog
// Add "Start" to message. (shorthand for M("Start"))
func (t *EzLog) TxtStart() *EzLog
// Add "OK"/"Fail" to message.
func (t *EzLog) Ok(data bool) *EzLog
// Add "Success"/"Fail" to message.
func (t *EzLog) Success(data bool) *EzLog
// Add "Yes"/"No" to message.
func (t *EzLog) YesNo(data bool) *EzLog
```

### Package Functions

```go
func New() *EzLog
func Clear() *EzLog
func Dump() *EzLog
// Get log level
func GetLogLevel() EzLogLevel
// Get log level prefix enable or not
func GetLogLevelPrefix() bool
// Set log level
func SetLogLevel(level EzLogLevel) *EzLog
// Enable/Disable log level prefix
func SetLogLevelPrefix(enable bool) *EzLog
// Set out function
func SetOutFunc(f OutFunc) *EzLog
// Set all log func to use fmt.Print()
func SetOutPrint() *EzLog
// Set all log func to use fmt.Println()
func SetOutPrintLn() *EzLog
// Enable/Disable trim on message
func SetTrim(enable bool) *EzLog
// Enable/Disable trim on `data`
func SetSkipEmpty(enable bool) *EzLog
// -- Set log message level
// Log message as level
func LogL(level EzLogLevel) *EzLog
// Log message without log level
func Log() *EzLog
// Log message as `EMERG`
func Emerg() *EzLog
// Log message as `ALERT`
func Alert() *EzLog
// Log message as `CRIT`
func Crit() *EzLog
// Log message as `ERR`
func Err() *EzLog
// Log message as `WARNING`
func Warning() *EzLog
// Log message as `NOTICE`
func Notice() *EzLog
// Log message as `INFO`
func Info() *EzLog
// Log message as `DEBUG`
func Debug() *EzLog
// Log message as `TRACE`
func Trace() *EzLog
// --- Output
func Out() *EzLog
func String() string
func StringP() *string
// --- Build log message
// enable/disable message log level prefix.
func Lp(enable bool) *EzLog
// enable/disable message log level prefix. (Alias for Lp())
func LogPrefix(enable bool) *EzLog
// Skip current message if `Msg` is empty. Current msg only. (shorthand of MsgSkipEmpty())
func Se() *EzLog
// Skip current message if `Msg` is empty. Current msg only.
func SkipEmpty() *EzLog
// --- base logging functions
// Append character/rune to message
func C(ch rune) *EzLog
// Add newline to message.
func L() *EzLog
// Add msg to log
func M(date any) *EzLog
// Add : after data
func N(data any) *EzLog
// --- Shorthand
// Add new line before data (shorthand for L().M())
func Lm(data any) *EzLog
// Add new line after data (shorthand for M().L())
func Ml(data any) *EzLog
// Add new line before data and : after (shorthand for L().N())
func Ln(data any) *EzLog
// Add : and new line after data (shorthand for N().L())
func Nl(data any) *EzLog
// --- Expressive func name
// Add new line to log (alias of L())
func NewLine() *EzLog
// Add msg to log (alias of M())
func Msg(data any) *EzLog
// Add : after data (alias of N())
func Name(data any) *EzLog
// --- Expressive alias
// Add new line before data (alias of L().M())
func NewLineMsg(data any) *EzLog
// Add new line after data (alias of M().L())
func MsgNewLine(data any) *EzLog
// Add new line before data and : after (alias of L().N())
func NewLineName(data any) *EzLog
// Add : and new line after data (alias of N().L())
func NameNewLine(data any) *EzLog
// -- Other shorthand
// Add tab to message.
func T() *EzLog
// Add tab to message
func Tab() *EzLog
// Add "End" to message. (shorthand for M("End"))
func TxtEnd() *EzLog
// Add "Start" to message. (shorthand for M("Start"))
func TxtStart() *EzLog
// Add "OK"/"Fail" to message.
func Ok(data bool) *EzLog
// Add "Success"/"Fail" to message.
func Success(data bool) *EzLog
// Add "Yes"/"No" to message.
func YesNo(data bool) *EzLog
```

## Example

Full example in top level example folder.

```go
func main() {
  var (
    log         = ezlog.New().SetLogLevel(ezlog.DebugLevel)
    N           = new(NUM).New()
    f32 float32 = 100.000001
    f64 float64 = 100.000001
    str string
  )

  fmt.Println("--- ezlog")
  log.Log().
    MsgNewLine(true).
    MsgNewLine(int16(-9910)).
    Name("0.008").MsgNewLine(float32(0.008)).
    Name("&f32").MsgNewLine(&f32).
    Name("&f64").MsgNewLine(&f64).
    Name("uint64").MsgNewLine(uint64(199999999999)).
    Name("N").L().MsgNewLine(N).
    Name("&N").L().Msg(&N).
    Out()

  str = log.String()

  fmt.Println("--- println")
  fmt.Println(str)
}
```

## License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
