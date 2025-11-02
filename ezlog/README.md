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

## Design

A simple log building library with only 4 base building functions. `Name` and `Message` part accept `any` as parameter. The design goal is ease of use.

## Installation

```sh
go get github.com/J-Siu/go-helper/v2
```

## Usage

```go
import "github.com/J-Siu/go-helper/v2/ezlog"
```

### Building a log message

A log message start with `ezlog` or an `Ezlog` instance. Follow by log message level, then building functions. Finally `Out()`.

```go
# use package level
# ezlog.<log level>.<building functions ...>.Out
ezlog.Debug.N("Key").M(value).Out()
```

```go
# use an ezlog instance
log := ezlog.New()
log.Debug.N("Key").M(value).Out()
```

### Log Control

Global|Per Message|Package level equivalent|Description
--|--|--|--
GetLogLevel/SetLogLevel(level)|n/a|yes|Set and get maximum log level a message will be printed (default: ERR)
SetLogLevelPrefix(bool)|Lp(bool)/LogPrefix(bool)|yes|enable/disable printing log level at beginning of log (default: true)
SetTrim(bool)|Tr(bool)/Trim(bool)|yes|Enable/disable trimming of name and message (default: true)
SetSkipEmpty|Se(bool)/SkipEmpty(bool)|yes|Do not log if message part is empty (default: false)

Per message settings should be used between message level setter and `Out()`.

### Log Message Level Setter

This needed to be set per log message and should be the first call in the call chain.

Function|Message Log Level|Numeric Level
--|--|--
LogL(EzLogLevel)|Programmatically set message log level|n/a
Log()|Log regardless of maximum log level|-2
Emerg()|EMERG|0
Alert()|ALERT|1
Crit()|CRIT|2
Err()|ERR|3
Warning()|WARNING|4
Notice()|NOTICE|5
Info()|INFO|6
Debug()|DEBUG|7
Trace()|TRACE|8

### Log Building Functions

`ezlog` design around 4 base log building functions. Shorthands functions are built on top of the above 4 functions. Each `M()` and `N()` are automatically separated by single space.

Function|Descriptive Alias|Count as message|Description
--|--|--|--
C(ch rune)||no|Add character/rune to the log
L() *EzLog|NewLine()|no|Add a newline to the log
M(data any)|Msg(data any)|yes|Add `data` add a message part to the log.
N(data any)|Name(data any)|no|Add `data` as a name part to the log. A `:` is added to the end of `data`

Shorthand functions for multiline log.

Function|Descriptive Alias|Count as message|Description
--|--|--|--
Lm(data any)|NewLineMsg(data any)|yes|Start message on a newline
Ml(data any)|MsgNewLine(data any)|yes|Add a newline after message
Ln(data any)|NewLineName(data any)|no|Start name on a newline
Nl(data any)|NameNewLine(data any)|no|Add a newline after name

Other shorthand functions.

Function|Descriptive Alias|Count as message|Description
--|--|--|--
T()|Tab()|no|Add a tab to log
TxtEnd()||yes|Add the word "End" to log
TxtStart()||yes|Add the word "Start" to log
OK(data bool)||yes|Add "OK"/"Fail" to log
Success(data bool)||yes|Add "Success"/"Fail" to log
YesNo(data bool)||yes|Add "Yes"/"No" to log

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
