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

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/J-Siu/go-helper/v2/strany"
)

type Level int8

// log level
const (
	LogLevel Level = iota - 2 // `LogLevel` is not exactly a log level. It is for logging regardless of log level
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
	StrAny         *strany.StrAny `json:"str_any,omitempty"`
	logLevel       Level
	logLevelPrefix bool
	msgLogLevel    Level
	outFunc        OutFunc
	strBuf         []string
	trim           bool
}

func (ez *ezlog) New() *ezlog {
	ez.SetLogLevel(ErrLevel)
	ez.SetLogLevelPrefix(true)
	ez.SetOutPrintLn()
	ez.SetTrim(true)
	ez.StrAny = new(strany.StrAny).New()
	return ez
}

// Clear message
func (ez *ezlog) Clear() *ezlog {
	ez.strBuf = nil
	return ez
}

// Get log level
func (ez *ezlog) GetLogLevel() Level { return ez.logLevel }

// Get log level prefix enable or not
func (ez *ezlog) GetLogLevelPrefix() bool { return ez.logLevelPrefix }

// Set out function
func (ez *ezlog) SetOutFunc(f OutFunc) *ezlog {
	ez.outFunc = f
	return ez
}

// Set out function to fmt.Print()
func (ez *ezlog) SetOutPrint() *ezlog {
	ez.SetOutFunc(func(str *string) { fmt.Print(*str) })
	return ez
}

// Set out function to fmt.Println()
func (ez *ezlog) SetOutPrintLn() *ezlog {
	ez.SetOutFunc(func(str *string) { fmt.Println(*str) })
	return ez
}

// Set log level
func (ez *ezlog) SetLogLevel(level Level) *ezlog {
	ez.logLevel = level
	return ez
}

// Set log level prefix true/false
func (ez *ezlog) SetLogLevelPrefix(enable bool) *ezlog {
	ez.logLevelPrefix = enable
	return ez
}

// Enable/Disable trim on `data`
func (ez *ezlog) SetTrim(enable bool) *ezlog {
	ez.trim = enable
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
	var strOut string
	if ez.msgLogLevel <= ez.logLevel {
		if ez.strBuf != nil {
			// str = strings.Join(l.strBuf, " ")
			for _, s := range ez.strBuf {
				_, size := utf8.DecodeLastRuneInString(strOut)
				if size > 0 && strOut[len(strOut)-size] != '\n' {
					strOut += " "
				}
				strOut += s
			}
		}
	}
	return &strOut
}

// --- Build log message

// Add newline to message. (shorthand for Ln())
func (ez *ezlog) L() *ezlog { return ez.Ln() }

// Add newline to message
func (ez *ezlog) Ln() *ezlog { return ez.Sp('\n') }

// Add msg to log
func (ez *ezlog) M(date any) *ezlog { return ez.Msg(date) }

// Add new line to message (shorthand for MsgLn())
func (ez *ezlog) Mn(date any) *ezlog { return ez.MsgLn(date) }

// Add new line to message (shorthand for MsgLn())
func (ez *ezlog) MLn(date any) *ezlog { return ez.MsgLn(date) }

// Add msg to log
func (ez *ezlog) Msg(data any) *ezlog {
	if ez.msgLogLevel <= ez.logLevel {
		tmp := *ez.StrAny.Str(data)
		if ez.trim {
			tmp = strings.Trim(tmp, "\n")
			tmp = strings.TrimSpace(tmp)
		} else {
		}
		ez.strBuf = append(ez.strBuf, tmp)
	}
	return ez
}

// Add new line to message (shorthand for Msg().Ln())
func (ez *ezlog) MsgLn(data any) *ezlog { return ez.Msg(data).Ln() }

// Add : after data (shorthand for Name())
func (ez *ezlog) N(data any) *ezlog { return ez.Name(data) }

// Add : and newline after data (shorthand for NameLn))
func (ez *ezlog) Nn(data any) *ezlog { return ez.NameLn(data) }

// Add : and newline after data (shorthand for NameLn))
func (ez *ezlog) NLn(data any) *ezlog { return ez.NameLn(data) }

// Add : after data (shorthand for Msg().Sp(':'))
func (ez *ezlog) Name(data any) *ezlog { return ez.Msg(data).Sp(':') }

// Add : and newline after data (shorthand for Msg().Sp(':').Ln())
func (ez *ezlog) NameLn(data any) *ezlog { return ez.Name(data).Ln() }

// Append character/rune to message (shorthand for Sp())
func (ez *ezlog) S(ch rune) *ezlog { return ez.Sp(ch) }

// Append character/rune to message
func (ez *ezlog) Sp(ch rune) *ezlog {
	if ez.msgLogLevel <= ez.logLevel {
		ez.strBuf[len(ez.strBuf)-1] += string(ch)
	}
	return ez
}

// Add tab to message. (shorthand for Tab())
func (ez *ezlog) T() *ezlog { return ez.Tab() }

// Add tab to message
func (ez *ezlog) Tab() *ezlog { return ez.Sp('\t') }

// Add "End" to message. (shorthand for Msg("End"))
func (ez *ezlog) TxtEnd() *ezlog { return ez.Msg("End") }

// Add "Start" to message. (shorthand for Msg("Start"))
func (ez *ezlog) TxtStart() *ezlog { return ez.Msg("Start") }

// -- Set log message level

func (ez *ezlog) Log() *ezlog {
	ez.Clear().msgLogLevel = LogLevel
	return ez
}
func (ez *ezlog) Emerg() *ezlog {
	ez.Clear().msgLogLevel = EmergLevel
	if ez.logLevelPrefix {
		ez.Name("EMERG")
	}
	return ez
}
func (ez *ezlog) Alert() *ezlog {
	ez.Clear().msgLogLevel = AlertLevel
	if ez.logLevelPrefix {
		ez.Name("ALERT")
	}
	return ez
}
func (ez *ezlog) Crit() *ezlog {
	ez.Clear().msgLogLevel = CritLevel
	if ez.logLevelPrefix {
		ez.Name("CRIT")
	}
	return ez
}
func (ez *ezlog) Err() *ezlog {
	ez.Clear().msgLogLevel = ErrLevel
	if ez.logLevelPrefix {
		ez.Name("ERR")
	}
	return ez
}
func (ez *ezlog) Warning() *ezlog {
	ez.Clear().msgLogLevel = WarningLevel
	if ez.logLevelPrefix {
		ez.Name("WARNING")
	}
	return ez
}
func (ez *ezlog) Notice() *ezlog {
	ez.Clear().msgLogLevel = NoticeLevel
	if ez.logLevelPrefix {
		ez.Name("NOTICE")
	}
	return ez
}
func (ez *ezlog) Info() *ezlog {
	ez.Clear().msgLogLevel = InfoLevel
	if ez.logLevelPrefix {
		ez.Name("INFO")
	}
	return ez
}
func (ez *ezlog) Debug() *ezlog {
	ez.Clear().msgLogLevel = DebugLevel
	if ez.logLevelPrefix {
		ez.Name("DEBUG")
	}
	return ez
}
func (ez *ezlog) Trace() *ezlog {
	ez.Clear().msgLogLevel = TraceLevel
	if ez.logLevelPrefix {
		ez.Name("TRACE")
	}
	return ez
}

// ---

var logger = New()

func New() *ezlog { return new(ezlog).New() }

// Get log level
func GetLogLevel() Level { return logger.GetLogLevel() }

// Get log level prefix enable or not
func GetLogLevelPrefix() bool { return logger.GetLogLevelPrefix() }

// Set log level
func SetLogLevel(level Level) *ezlog { return logger.SetLogLevel(level) }

// Enable/Disable log level prefix
func SetLogLevelPrefix(enable bool) *ezlog { return logger.SetLogLevelPrefix(enable) }

// Set all log func to use fmt.Print()
func SetOutPrint() *ezlog { return logger.SetOutPrint() }

// Set all log func to use fmt.Println()
func SetOutPrintLn() *ezlog { return logger.SetOutPrintLn() }

// Enable/Disable trim on message
func SetTrim(enable bool) *ezlog { return logger.SetTrim(enable) }

func String() string   { return logger.String() }
func StringP() *string { return logger.StringP() }

func Log() *ezlog     { return logger.Log() }
func Emerg() *ezlog   { return logger.Emerg() }
func Alert() *ezlog   { return logger.Alert() }
func Crit() *ezlog    { return logger.Crit() }
func Err() *ezlog     { return logger.Err() }
func Warning() *ezlog { return logger.Warning() }
func Notice() *ezlog  { return logger.Notice() }
func Info() *ezlog    { return logger.Info() }
func Debug() *ezlog   { return logger.Debug() }
func Trace() *ezlog   { return logger.Trace() }
