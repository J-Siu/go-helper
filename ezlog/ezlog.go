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

type OutFunc func(msg *string)

type EzLog struct {
	StrAny         *strany.StrAny `json:"str_any,omitempty"`
	logLevel       Level
	logLevelPrefix bool
	msgLogLevel    Level
	outFunc        OutFunc
	strBuf         []string
	trim           bool
}

func (ez *EzLog) New() *EzLog {
	ez.SetLogLevel(ERR)
	ez.SetLogLevelPrefix(true)
	ez.SetOutPrintLn()
	ez.SetTrim(true)
	ez.StrAny = new(strany.StrAny).New()
	return ez
}

// Clear message
func (ez *EzLog) Clear() *EzLog {
	ez.strBuf = nil
	return ez
}

// Get log level
func (ez *EzLog) GetLogLevel() Level { return ez.logLevel }

// Get log level prefix enable or not
func (ez *EzLog) GetLogLevelPrefix() bool { return ez.logLevelPrefix }

// Set out function
func (ez *EzLog) SetOutFunc(f OutFunc) *EzLog {
	ez.outFunc = f
	return ez
}

// Set out function to fmt.Print()
func (ez *EzLog) SetOutPrint() *EzLog {
	ez.SetOutFunc(func(str *string) { fmt.Print(*str) })
	return ez
}

// Set out function to fmt.Println()
func (ez *EzLog) SetOutPrintLn() *EzLog {
	ez.SetOutFunc(func(str *string) { fmt.Println(*str) })
	return ez
}

// Set log level
func (ez *EzLog) SetLogLevel(level Level) *EzLog {
	ez.logLevel = level
	return ez
}

// Set log level prefix true/false
func (ez *EzLog) SetLogLevelPrefix(enable bool) *EzLog {
	ez.logLevelPrefix = enable
	return ez
}

// Enable/Disable trim on `data`
func (ez *EzLog) SetTrim(enable bool) *EzLog {
	ez.trim = enable
	return ez
}

// --- Output

func (ez *EzLog) Out() *EzLog {
	if ez.msgLogLevel <= ez.logLevel {
		ez.outFunc(ez.StringP())
	}
	return ez
}

func (ez *EzLog) String() string { return *ez.StringP() }

func (ez *EzLog) StringP() *string {
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
func (ez *EzLog) L() *EzLog { return ez.Ln() }

// Add newline to message
func (ez *EzLog) Ln() *EzLog { return ez.Sp('\n') }

// Add msg to log
func (ez *EzLog) M(date any) *EzLog { return ez.Msg(date) }

// Add new line to message (shorthand for MsgLn())
func (ez *EzLog) Mn(date any) *EzLog { return ez.MsgLn(date) }

// Add new line to message (shorthand for MsgLn())
func (ez *EzLog) MLn(date any) *EzLog { return ez.MsgLn(date) }

// Add msg to log
func (ez *EzLog) Msg(data any) *EzLog {
	if ez.msgLogLevel <= ez.logLevel {
		tmp := *ez.StrAny.Any(data)
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
func (ez *EzLog) MsgLn(data any) *EzLog { return ez.Msg(data).Ln() }

// Add : after data (shorthand for Name())
func (ez *EzLog) N(data any) *EzLog { return ez.Name(data) }

// Add : and newline after data (shorthand for NameLn))
func (ez *EzLog) Nn(data any) *EzLog { return ez.NameLn(data) }

// Add : and newline after data (shorthand for NameLn))
func (ez *EzLog) NLn(data any) *EzLog { return ez.NameLn(data) }

// Add : after data (shorthand for Msg().Sp(':'))
func (ez *EzLog) Name(data any) *EzLog { return ez.Msg(data).Sp(':') }

// Add : and newline after data (shorthand for Msg().Sp(':').Ln())
func (ez *EzLog) NameLn(data any) *EzLog { return ez.Name(data).Ln() }

// Append character/rune to message (shorthand for Sp())
func (ez *EzLog) S(ch rune) *EzLog { return ez.Sp(ch) }

// Append character/rune to message
func (ez *EzLog) Sp(ch rune) *EzLog {
	if ez.msgLogLevel <= ez.logLevel {
		count := len(ez.strBuf)
		if count == 0 {
			ez.strBuf = append(ez.strBuf, string(ch))
		} else {
			ez.strBuf[count-1] += string(ch)
		}
	}
	return ez
}

// Add tab to message. (shorthand for Tab())
func (ez *EzLog) T() *EzLog { return ez.Tab() }

// Add tab to message
func (ez *EzLog) Tab() *EzLog { return ez.Sp('\t') }

// Add "End" to message. (shorthand for Msg("End"))
func (ez *EzLog) TxtEnd() *EzLog { return ez.Msg("End") }

// Add "Start" to message. (shorthand for Msg("Start"))
func (ez *EzLog) TxtStart() *EzLog { return ez.Msg("Start") }

// -- Set log message level

// Log message as `level`
func (ez *EzLog) LogL(level Level) *EzLog {
	ez.Clear().msgLogLevel = level
	return ez
}

// Log message without level (no level prefix)
func (ez *EzLog) Log() *EzLog {
	ez.Clear().msgLogLevel = LogLevel
	return ez
}

// Log message as `EMERG`
func (ez *EzLog) Emerg() *EzLog {
	ez.Clear().msgLogLevel = EMERG
	if ez.logLevelPrefix {
		ez.Name(ez.msgLogLevel.String())
	}
	return ez
}

// Log message as `ALERT`
func (ez *EzLog) Alert() *EzLog {
	ez.Clear().msgLogLevel = ALERT
	if ez.logLevelPrefix {
		ez.Name(ez.msgLogLevel.String())
	}
	return ez
}

// Log message as `CRIT`
func (ez *EzLog) Crit() *EzLog {
	ez.Clear().msgLogLevel = CRIT
	if ez.logLevelPrefix {
		ez.Name(ez.msgLogLevel.String())
	}
	return ez
}

// Log message as `ERR`
func (ez *EzLog) Err() *EzLog {
	ez.Clear().msgLogLevel = ERR
	if ez.logLevelPrefix {
		ez.Name(ez.msgLogLevel.String())
	}
	return ez
}

// Log message as `WARN`
func (ez *EzLog) Warning() *EzLog {
	ez.Clear().msgLogLevel = WARNING
	if ez.logLevelPrefix {
		ez.Name(ez.msgLogLevel.String())
	}
	return ez
}

// Log message as `NOTICE`
func (ez *EzLog) Notice() *EzLog {
	ez.Clear().msgLogLevel = NOTICE
	if ez.logLevelPrefix {
		ez.Name(ez.msgLogLevel.String())
	}
	return ez
}

// Log message as `INFO`
func (ez *EzLog) Info() *EzLog {
	ez.Clear().msgLogLevel = INFO
	if ez.logLevelPrefix {
		ez.Name(ez.msgLogLevel.String())
	}
	return ez
}

// Log message as `DEBUG`
func (ez *EzLog) Debug() *EzLog {
	ez.Clear().msgLogLevel = DEBUG
	if ez.logLevelPrefix {
		ez.Name(ez.msgLogLevel.String())
	}
	return ez
}

// Log message as `TRACE`
func (ez *EzLog) Trace() *EzLog {
	ez.Clear().msgLogLevel = TRACE
	if ez.logLevelPrefix {
		ez.Name(ez.msgLogLevel.String())
	}
	return ez
}

// ---

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
