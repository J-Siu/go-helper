# ezlog

A simple log module with Linux log level:

- (-2) LogLevel // Not exactly a log level. It is for logging regardless of log level
- (-1) Disabled
0. Emerg
1. Alert
2. Crit
3. Err
4. Warning
5. Notice
6. Info
7. Debug
8. Trace

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
  LogLevel Level = iota - 2 // `LogLevel` is not exactly a log level. It is for logging regardless of log level
  Disabled
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
func (ez *EzLog) New() *EzLog
func (ez *EzLog) Clear() *EzLog
func (ez *EzLog) GetLogLevel() Level
func (ez *EzLog) GetLogLevelPrefix() bool
func (ez *EzLog) SetOutFunc(f OutFunc) *EzLog
func (ez *EzLog) SetOutPrint() *EzLog
func (ez *EzLog) SetOutPrintLn() *EzLog
func (ez *EzLog) SetLogLevel(level Level) *EzLog
func (ez *EzLog) SetLogLevelPrefix(enable bool) *EzLog
func (ez *EzLog) SetTrim(enable bool) *EzLog
func (ez *EzLog) Out() *EzLog
func (ez *EzLog) String() string
func (ez *EzLog) StringP() *string
func (ez *EzLog) L() *EzLog
func (ez *EzLog) Ln() *EzLog
func (ez *EzLog) M(date any) *EzLog
func (ez *EzLog) Mn(date any) *EzLog
func (ez *EzLog) MLn(date any) *EzLog
func (ez *EzLog) Msg(data any) *EzLog
func (ez *EzLog) MsgLn(data any) *EzLog
func (ez *EzLog) N(data any) *EzLog
func (ez *EzLog) Nn(data any) *EzLog
func (ez *EzLog) NLn(data any) *EzLog
func (ez *EzLog) Name(data any) *EzLog
func (ez *EzLog) NameLn(data any) *EzLog
func (ez *EzLog) S(ch rune) *EzLog
func (ez *EzLog) Sp(ch rune) *EzLog
func (ez *EzLog) T() *EzLog
func (ez *EzLog) Tab() *EzLog
func (ez *EzLog) TxtEnd() *EzLog
func (ez *EzLog) TxtStart() *EzLog
func (ez *EzLog) LogL(level Level) *EzLog
func (ez *EzLog) Log() *EzLog
func (ez *EzLog) Emerg() *EzLog
func (ez *EzLog) Alert() *EzLog
func (ez *EzLog) Crit() *EzLog
func (ez *EzLog) Err() *EzLog
func (ez *EzLog) Warning() *EzLog
func (ez *EzLog) Notice() *EzLog
func (ez *EzLog) Info() *EzLog
func (ez *EzLog) Debug() *EzLog
func (ez *EzLog) Trace() *EzLog
```

### Package Functions

```go
func New() *EzLog
func GetLogLevel() Level
func GetLogLevelPrefix() bool
func SetLogLevel(level Level) *EzLog
func SetLogLevelPrefix(enable bool) *EzLog
func SetOutPrint() *EzLog
func SetOutPrintLn() *EzLog
func SetTrim(enable bool) *EzLog
func String() string
func StringP() *string
func LogL(level Level) *EzLog
func Log() *EzLog
func Emerg() *EzLog
func Alert() *EzLog
func Crit() *EzLog
func Err() *EzLog
func Warning() *EzLog
func Notice() *EzLog
func Info() *EzLog
func Debug() *EzLog
func Trace() *EzLog
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
    MsgLn(true).
    MsgLn(int16(-9910)).
    Name("0.008").MsgLn(float32(0.008)).
    Name("&f32").MsgLn(&f32).
    Name("&f64").MsgLn(&f64).
    MsgLn(uint64(199999999999)).
    Name("N").Ln().MsgLn(N).
    Name("&N").Ln().Msg(&N).
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
