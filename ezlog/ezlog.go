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
//  0. Disable)
//  1. Emerg
//  2. Alert
//  3. Crit
//  4. Err
//  5. Warning
//  6. Notice
//  7. Info
//  8. Debug
//  9. Trace
package ezlog

import (
	"fmt"
	"unicode/utf8"

	"github.com/J-Siu/go-helper/v2/str"
)

type Level int8

// log level
const (
	LogLevel Level = iota - 2 // Not exactly a log level. It is for logging regardless of log level
	Disabled
	EmergLevel
	AlertLevel
	CritLevel
	ErrLevel
	WarningLevel
	NoticeLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

type OutFunc func(msg *string)

type ezlog struct {
	StrAny      *str.Any
	logLevel    Level
	msgLogLevel Level
	outFunc     OutFunc
	strBuf      []string
}

func (ez *ezlog) New() *ezlog {
	ez.StrAny = new(str.Any).New().IndentEnable(true)
	ez.SetLogLevel(ErrLevel)
	ez.SetOutPrintLn()
	return ez
}

func (ez *ezlog) SetOutFunc(f OutFunc) *ezlog {
	ez.outFunc = f
	return ez
}

// Set out
func (ez *ezlog) SetOutPrint() *ezlog {
	ez.SetOutFunc(func(str *string) { fmt.Print(*str) })
	return ez
}

// Set out
func (ez *ezlog) SetOutPrintLn() *ezlog {
	ez.SetOutFunc(func(str *string) { fmt.Println(*str) })
	return ez
}

// Set log level
func (ez *ezlog) SetLogLevel(level Level) *ezlog {
	ez.logLevel = level
	return ez
}

// Get log level
func (ez *ezlog) GetLogLevel() Level { return ez.logLevel }

// Clear message
func (ez *ezlog) Clear() *ezlog {
	ez.strBuf = nil
	return ez
}

// --- Output

func (ez *ezlog) Out() *ezlog {
	if ez.msgLogLevel <= ez.logLevel {
		ez.outFunc(ez.StringP())
	}
	return ez
}

func (ez *ezlog) String() string { return *ez.StringP() }

func (ez *ezlog) StringP() *string {
	str := ""
	if ez.msgLogLevel <= ez.logLevel {
		if ez.strBuf != nil {
			// str = strings.Join(l.strBuf, " ")
			for _, s := range ez.strBuf {
				_, size := utf8.DecodeLastRuneInString(str)
				if size > 0 && str[len(str)-size] != '\n' {
					str += " "
				}
				str += s
			}
		}
	}
	return &str
}

// --- Build log message

// Add msg to log
func (ez *ezlog) Msg(data any) *ezlog {
	if ez.msgLogLevel <= ez.logLevel {
		ez.strBuf = append(ez.strBuf, ez.StrAny.Str(data))
	}
	return ez
}

// Add separator to message
func (ez *ezlog) Sp(data any) *ezlog {
	if ez.msgLogLevel <= ez.logLevel {
		ez.strBuf[len(ez.strBuf)-1] = ez.strBuf[len(ez.strBuf)-1] + ez.StrAny.Str(data)
	}
	return ez
}

// Add newline to message
func (ez *ezlog) Ln() *ezlog { return ez.Msg("\n") }

// Log on new line
func (ez *ezlog) MsgLn(data any) *ezlog { return ez.Msg(data).Ln() }

// Add : after data
func (ez *ezlog) Name(data any) *ezlog { return ez.Msg(data).Sp(":") }

// Add : and newline after data
func (ez *ezlog) NameLn(data any) *ezlog { return ez.Name(data).Ln() }

// -- Set log message level

func (ez *ezlog) Log() *ezlog {
	ez.Clear().msgLogLevel = LogLevel
	return ez
}
func (ez *ezlog) Emerg() *ezlog {
	ez.Clear().msgLogLevel = EmergLevel
	return ez
}
func (ez *ezlog) Alert() *ezlog {
	ez.Clear().msgLogLevel = AlertLevel
	return ez
}
func (ez *ezlog) Crit() *ezlog {
	ez.Clear().msgLogLevel = CritLevel
	return ez
}
func (ez *ezlog) Err() *ezlog {
	ez.Clear().msgLogLevel = ErrLevel
	return ez
}
func (ez *ezlog) Warning() *ezlog {
	ez.Clear().msgLogLevel = WarningLevel
	return ez
}
func (ez *ezlog) Notice() *ezlog {
	ez.Clear().msgLogLevel = NoticeLevel
	return ez
}
func (ez *ezlog) Info() *ezlog {
	ez.Clear().msgLogLevel = InfoLevel
	return ez
}
func (ez *ezlog) Debug() *ezlog {
	ez.Clear().msgLogLevel = DebugLevel
	return ez
}
func (ez *ezlog) Trace() *ezlog {
	ez.Clear().msgLogLevel = TraceLevel
	return ez
}

// ---

var log = New()

func New() *ezlog {
	return new(ezlog).New()
}

// Set all log func to use fmt.Print()
func SetOutPrint() *ezlog { return log.SetOutPrint() }

// Set all log func to use fmt.Println()
func SetOutPrintLn() *ezlog { return log.SetOutPrintLn() }

// Get log level
func GetLogLevel() Level { return log.GetLogLevel() }

// Set log level
func SetLogLevel(level Level) *ezlog { return log.SetLogLevel(level) }

func Log() *ezlog {
	log.Clear().msgLogLevel = LogLevel
	return log
}
func Emerg() *ezlog {
	log.Clear().msgLogLevel = EmergLevel
	return log
}
func Alert() *ezlog {
	log.Clear().msgLogLevel = AlertLevel
	return log
}
func Crit() *ezlog {
	log.Clear().msgLogLevel = CritLevel
	return log
}
func Err() *ezlog {
	log.Clear().msgLogLevel = ErrLevel
	return log
}
func Warning() *ezlog {
	log.Clear().msgLogLevel = WarningLevel
	return log
}
func Notice() *ezlog {
	log.Clear().msgLogLevel = NoticeLevel
	return log
}
func Info() *ezlog {
	log.Clear().msgLogLevel = InfoLevel
	return log
}
func Debug() *ezlog {
	log.Clear().msgLogLevel = DebugLevel
	return log
}
func Trace() *ezlog {
	log.Clear().msgLogLevel = TraceLevel
	return log
}
