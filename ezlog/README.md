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

func (t *EzLog) New() *EzLog
func (t *EzLog) Clear() *EzLog
func (t *EzLog) GetLogLevel() EzLogLevel
func (t *EzLog) GetLogLevelPrefix() bool
func (t *EzLog) SetLogLevel(level EzLogLevel) *EzLog
func (t *EzLog) SetLogLevelPrefix(enable bool) *EzLog
func (t *EzLog) SetOutFunc(f OutFunc) *EzLog
func (t *EzLog) SetOutPrint() *EzLog
func (t *EzLog) SetOutPrintLn() *EzLog
func (t *EzLog) SetTrim(enable bool) *EzLog
func (t *EzLog) SetSkipEmpty(enable bool) *EzLog
func (t *EzLog) LogL(level EzLogLevel) *EzLog
func (t *EzLog) Log() *EzLog
func (t *EzLog) Emerg() *EzLog
func (t *EzLog) Alert() *EzLog
func (t *EzLog) Crit() *EzLog
func (t *EzLog) Err() *EzLog
func (t *EzLog) Warning() *EzLog
func (t *EzLog) Notice() *EzLog
func (t *EzLog) Info() *EzLog
func (t *EzLog) Debug() *EzLog
func (t *EzLog) Trace() *EzLog
func (t *EzLog) Dump() *EzLog
func (t *EzLog) Out() *EzLog
func (t *EzLog) String() string
func (t *EzLog) StringP() *string
func (t *EzLog) build(data any, isMsg bool) *EzLog
func (t *EzLog) Lp(enable bool) *EzLog
func (t *EzLog) LogPrefix(enable bool) *EzLog
func (t *EzLog) Se() *EzLog
func (t *EzLog) SkipEmpty() *EzLog
func (t *EzLog) C(ch rune) *EzLog
func (t *EzLog) L() *EzLog
func (t *EzLog) M(data any) *EzLog
func (t *EzLog) Mn(data any) *EzLog
func (t *EzLog) Msg(data any) *EzLog
func (t *EzLog) MsgLn(data any) *EzLog
func (t *EzLog) N(data any) *EzLog
func (t *EzLog) Nn(data any) *EzLog
func (t *EzLog) Name(data any) *EzLog
func (t *EzLog) NameLn(data any) *EzLog
func (t *EzLog) T() *EzLog
func (t *EzLog) Tab() *EzLog
func (t *EzLog) TxtEnd() *EzLog
func (t *EzLog) TxtStart() *EzLog
func (t *EzLog) Ok(data bool) *EzLog
func (t *EzLog) Success(data bool) *EzLog
func (t *EzLog) YesNo(data bool) *EzLog
```

### Package Functions

```go
func New() *EzLog
func Clear() *EzLog
func Dump() *EzLog
func GetLogLevel() EzLogLevel
func GetLogLevelPrefix() bool
func SetLogLevel(level EzLogLevel) *EzLog
func SetLogLevelPrefix(enable bool) *EzLog
func SetOutFunc(f OutFunc) *EzLog
func SetOutPrint() *EzLog
func SetOutPrintLn() *EzLog
func SetTrim(enable bool) *EzLog
func SetSkipEmpty(enable bool) *EzLog
func LogL(level EzLogLevel) *EzLog
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
func Out() *EzLog
func String() string
func StringP() *string
func Lp(enable bool) *EzLog
func LogPrefix(enable bool) *EzLog
func Se() *EzLog
func SkipEmpty() *EzLog
func C(ch rune) *EzLog
func L() *EzLog
func M(date any) *EzLog
func Mn(date any) *EzLog
func Msg(data any) *EzLog
func MsgLn(data any) *EzLog
func N(data any) *EzLog
func Nn(data any) *EzLog
func Name(data any) *EzLog
func NameLn(data any) *EzLog
func T() *EzLog
func Tab() *EzLog
func TxtEnd() *EzLog
func TxtStart() *EzLog
func Ok(data bool) *EzLog
func Success(data bool) *EzLog
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
